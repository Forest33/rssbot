// Package usecase provides business logic.
package usecase

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/microcosm-cc/bluemonday"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/database"
	"github.com/forest33/rssbot/pkg/logger"
	"github.com/forest33/rssbot/pkg/structs"
)

// BotUseCase object capable of interacting with BotUseCase
type BotUseCase struct {
	ctx               context.Context
	cfg               *entity.BotConfig
	log               *logger.Zerolog
	db                *database.Database
	feedsRepo         FeedsRepo
	usersRepo         UsersRepo
	subscriptionsRepo SubscriptionsRepo
	bot               *tgbotapi.BotAPI
	updatesChan       tgbotapi.UpdatesChannel
	commandChan       chan *commandJob
	itemsChan         chan *itemsJob
	senderChan        chan *senderJob
	sanitizer         *bluemonday.Policy
}

// UsersRepo is the common interface implemented UsersRepository methods
type UsersRepo interface {
	Create(ctx context.Context, in *entity.User) (*entity.User, error)
	Get(ctx context.Context, filter *entity.UsersFilter) ([]*entity.User, error)
}

// SubscriptionsRepo is the common interface implemented SubscriptionsRepository methods
type SubscriptionsRepo interface {
	GetByUserID(ctx context.Context, id int64) ([]*entity.Subscription, error)
	GetByFeedID(ctx context.Context, id string) ([]*entity.Subscription, error)
	Create(ctx context.Context, in *entity.Subscription) (*entity.Subscription, error)
	Delete(ctx context.Context, id string) error
}

type commandJob struct {
	msg *tgbotapi.Message
}

type senderJob struct {
	chatID   int64
	messages []string
}

type itemsJob struct {
	feed  *entity.Feed
	items []*entity.FeedItem
}

// NewBotUseCase creates a new BotUseCase
func NewBotUseCase(ctx context.Context, cfg *entity.BotConfig, log *logger.Zerolog, db *database.Database, feedsRepo FeedsRepo, usersRepo UsersRepo, subscriptionsRepo SubscriptionsRepo) (*BotUseCase, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	uc := &BotUseCase{
		ctx:               ctx,
		cfg:               cfg,
		log:               log,
		db:                db,
		feedsRepo:         feedsRepo,
		usersRepo:         usersRepo,
		subscriptionsRepo: subscriptionsRepo,
		commandChan:       make(chan *commandJob, cfg.CommandWorkersPoolSize),
		itemsChan:         make(chan *itemsJob, cfg.SenderWorkersPoolSize),
		senderChan:        make(chan *senderJob, cfg.SenderWorkersPoolSize),
	}

	return uc, uc.init()
}

func (uc *BotUseCase) Start() {
	go uc.loop()
}

func (uc *BotUseCase) Send(feed *entity.Feed, items []*entity.FeedItem) {
	uc.itemsChan <- &itemsJob{
		feed:  feed,
		items: items,
	}
}

func (uc *BotUseCase) init() error {
	var err error
	uc.bot, err = tgbotapi.NewBotAPI(uc.cfg.Token)
	if err != nil {
		return err
	}
	uc.bot.Debug = uc.cfg.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = uc.cfg.UpdateTimeout
	uc.updatesChan = uc.bot.GetUpdatesChan(u)

	uc.sanitizer = bluemonday.NewPolicy()
	uc.sanitizer.AllowStandardURLs()
	uc.sanitizer.AllowAttrs("href").OnElements("a")

	for i := 0; i < uc.cfg.CommandWorkersPoolSize; i++ {
		go uc.commandWorker()
	}

	for i := 0; i < uc.cfg.SenderWorkersPoolSize; i++ {
		go uc.itemsWorker()
		go uc.senderWorker()
	}

	return nil
}

func (uc *BotUseCase) loop() {
	defer func() {
		uc.bot.StopReceivingUpdates()
	}()

	for {
		select {
		case <-uc.ctx.Done():
			return
		case update, ok := <-uc.updatesChan:
			if !ok {
				return
			}
			if update.Message == nil || update.Message.Text == "" {
				continue
			}

			uc.commandChan <- &commandJob{
				msg: update.Message,
			}
		}
	}
}

func (uc *BotUseCase) commandWorker() {
	errFunc := func(msg *tgbotapi.Message, err error) {
		uc.log.Error().
			Err(err).
			Int64("user_id", msg.From.ID).
			Str("text", msg.Text).
			Msg("processing command error")
	}

	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.commandChan:
			if !ok {
				return
			}

			err := uc.commandHandler(job.msg)
			if err != nil {
				errFunc(job.msg, err)
				continue
			}

			uc.log.Debug().
				Int64("user_id", job.msg.From.ID).
				Str("text", job.msg.Text).
				Msg("processing command successful")
		}
	}
}

func (uc *BotUseCase) itemsWorker() {
	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.itemsChan:
			if !ok {
				return
			}
			if err := uc.sendItems(job); err == nil {
				if _, err := uc.feedsRepo.Update(uc.ctx, job.feed); err != nil {
					uc.log.Error().Err(err).Str("feed_id", job.feed.ID).Msg("failed to mark feed")
				}
			}
		}
	}
}

func (uc *BotUseCase) senderWorker() {
	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.senderChan:
			if !ok {
				return
			}
			sentCount := 0
			for _, text := range job.messages {
				sentCount++
				if err := uc.send(job.chatID, text); err != nil {
					uc.log.Error().Err(err).
						Int64("chat_id", job.chatID).
						Msg("failed to send message")
					continue
				}
			}
			uc.log.Debug().Int64("chat_id", job.chatID).Msgf("%d messages sent", sentCount)
		}
	}
}

func (uc *BotUseCase) commandHandler(msg *tgbotapi.Message) error {
	sendError := func(err error) {
		if err := uc.reply(msg, entity.GetErrorMessage(err)); err != nil {
			uc.log.Error().Err(err).Msg("failed to send error message")
		}
	}

	cmd, args, err := entity.ParseCommand(msg.Text)
	if err != nil {
		sendError(err)
		return err
	}

	ctx, err := uc.db.BeginTransaction(uc.ctx)
	if err != nil {
		uc.log.Error().Err(err).Msg("failed to begin transaction")
		return entity.ErrInternal
	}
	defer uc.db.CommitTransaction(ctx, err)

	switch cmd {
	case entity.CommandStart:
		err = uc.createUser(ctx, msg, args)
	case entity.CommandAddSubscription:
		err = uc.addSubscription(ctx, msg, args)
	case entity.CommandListSubscriptions:
		err = uc.getSubscriptions(ctx, msg, args)
	case entity.CommandDeleteSubscription:
		err = uc.deleteSubscription(ctx, msg, args)
	case entity.CommandHelp:
		err = uc.reply(msg, entity.MessageHelp)
	default:
		err = entity.ErrUnknownBotCommand
	}

	if err != nil {
		sendError(err)
	}

	return err
}

func (uc *BotUseCase) sendItems(job *itemsJob) error {
	subs, err := uc.subscriptionsRepo.GetByFeedID(uc.ctx, job.feed.ID)
	if err != nil {
		uc.log.Error().Err(err).Str("feed_id", job.feed.ID).Msg("failed to get feed subscriptions")
		return err
	}

	messages := uc.prepareMessages(job.items)

	for _, s := range subs {
		uc.senderChan <- &senderJob{
			chatID:   s.UserID,
			messages: messages,
		}
	}

	return nil
}

func (uc *BotUseCase) prepareMessages(items []*entity.FeedItem) []string {
	messages := make([]string, len(items))
	for i := len(items) - 1; i >= 0; i-- {
		msg := fmt.Sprintf("<b><a href=\"%s\">%s</a></b>\n%s", items[i].Link, items[i].Title, uc.sanitizer.Sanitize(items[i].Description))
		messages[i] = msg
	}
	return messages
}

func (uc *BotUseCase) getUserFeeds(ctx context.Context, userID int64) ([]*entity.Subscription, []*entity.Feed, error) {
	subs, err := uc.subscriptionsRepo.GetByUserID(ctx, userID)
	if err != nil {
		uc.log.Error().Err(err).
			Int64("user_id", userID).
			Msg("failed to get user subscriptions")
		return nil, nil, entity.ErrInternal
	}

	if len(subs) == 0 {
		return nil, nil, nil
	}

	feedIds := structs.Map(subs, func(s *entity.Subscription) string { return s.FeedID })

	feeds, err := uc.feedsRepo.Get(ctx, &entity.FeedFilter{ID: feedIds})
	if err != nil {
		uc.log.Error().Err(err).
			Int64("user_id", userID).
			Msg("failed to get feeds")
		return nil, nil, entity.ErrInternal
	}

	return subs, feeds, nil
}

func (uc *BotUseCase) reply(msg *tgbotapi.Message, text string) error {
	if len(text) == 0 {
		return nil
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	reply.ReplyToMessageID = msg.MessageID
	reply.DisableWebPagePreview = true

	_, err := uc.bot.Send(reply)
	if err != nil {
		uc.log.Error().Err(err).
			Int64("user_id", msg.From.ID).
			Str("text", text).
			Msg("failed to send reply message")
	}

	return err
}

func (uc *BotUseCase) send(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := uc.bot.Send(msg)

	return err
}
