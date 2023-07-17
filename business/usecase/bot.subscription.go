package usecase

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/structs"
)

func (uc *BotUseCase) addSubscription(ctx context.Context, msg *tgbotapi.Message, args []string) error {
	if len(args) != 1 {
		return entity.ErrWrongNumberOfArguments
	} else if entity.ValidateURL(args[0]) != nil {
		return entity.ErrWrongURL
	}

	feedURL, err := parserUseCase.CheckFeedURL(ctx, args[0])
	if err != nil {
		return err
	}

	feed, err := uc.feedsRepo.Create(ctx, &entity.Feed{
		SiteURL: args[0],
		FeedURL: feedURL,
	})
	if err != nil {
		uc.log.Error().Err(err).
			Int64("user_id", msg.From.ID).
			Str("site_url", args[0]).
			Str("feed_url", feedURL).
			Msg("failed to create feed")
		return entity.ErrInternal
	}

	_, err = uc.subscriptionsRepo.Create(ctx, &entity.Subscription{
		UserID: msg.From.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		uc.log.Error().Err(err).
			Int64("user_id", msg.From.ID).
			Str("feed_id", feed.ID).
			Msg("failed to add user subscription")
		return entity.ErrInternal
	}

	_ = uc.reply(msg, entity.MessageSubscriptionCreated)

	parserUseCase.LoadFeed(feed)

	return nil
}

func (uc *BotUseCase) getSubscriptions(ctx context.Context, msg *tgbotapi.Message, _ []string) error {
	_, feeds, err := uc.getUserFeeds(ctx, msg.From.ID)
	if err != nil {
		return err
	} else if feeds == nil {
		return uc.reply(msg, entity.MessageNoSubscriptions)
	}

	sb := strings.Builder{}
	for i, f := range feeds {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, structs.If(f.Title == "", f.SiteURL, f.Title)))
	}

	return uc.reply(msg, sb.String())
}

func (uc *BotUseCase) deleteSubscription(ctx context.Context, msg *tgbotapi.Message, args []string) error {
	if len(args) != 1 {
		return entity.ErrWrongNumberOfArguments
	}

	subNumber, err := entity.ValidateNumber(args[0])
	if err != nil {
		return err
	}

	subs, feeds, err := uc.getUserFeeds(ctx, msg.From.ID)
	if err != nil {
		return err
	} else if feeds == nil {
		return uc.reply(msg, entity.MessageNoSubscriptions)
	}

	if len(feeds) < subNumber {
		return entity.ErrWrongSubscriptionNumber
	}

	var subscriptionID string
	for _, s := range subs {
		if s.FeedID == feeds[subNumber-1].ID {
			subscriptionID = s.ID
			break
		}
	}

	if subscriptionID == "" {
		return entity.ErrInternal
	}

	if err := uc.subscriptionsRepo.Delete(ctx, subscriptionID); err != nil {
		uc.log.Error().Err(err).
			Str("subscription_id", subscriptionID).
			Msg("failed to delete subscription")
		return entity.ErrInternal
	}

	_ = uc.reply(msg, entity.MessageSubscriptionDeleted)

	if feedSubs, err := uc.subscriptionsRepo.GetByFeedID(ctx, feeds[subNumber-1].ID); err != nil {
		uc.log.Error().Err(err).
			Str("feed_id", feeds[subNumber-1].ID).
			Msg("failed to get feed subscriptions")
		return entity.ErrInternal
	} else if len(feedSubs) == 1 && feedSubs[0].UserID == msg.From.ID {
		if err := uc.feedsRepo.Delete(ctx, feeds[subNumber-1].ID); err != nil {
			uc.log.Error().Err(err).
				Str("feed_id", feeds[subNumber-1].ID).
				Msg("failed to delete feed")
			return entity.ErrInternal
		}
	}

	return nil
}
