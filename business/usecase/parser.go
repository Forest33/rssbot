// Package usecase provides business logic.
package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/logger"
	"github.com/forest33/rssbot/pkg/structs"
)

// ParserUseCase object capable of interacting with ParserUseCase
type ParserUseCase struct {
	ctx         context.Context
	cfg         *entity.ParserConfig
	log         *logger.Zerolog
	feedsRepo   FeedsRepo
	workersChan chan *feedJob
}

type feedJob struct {
	feed *entity.Feed
}

const (
	userAgent = "rssbot (https://github.com/forest33/rssbot)"
)

// NewParserUseCase creates a new ParserUseCase
func NewParserUseCase(ctx context.Context, cfg *entity.ParserConfig, log *logger.Zerolog, feedsRepo FeedsRepo) (*ParserUseCase, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	uc := &ParserUseCase{
		ctx:         ctx,
		cfg:         cfg,
		log:         log,
		feedsRepo:   feedsRepo,
		workersChan: make(chan *feedJob, cfg.WorkersPoolSize),
	}

	uc.init()

	return uc, nil
}

func (uc *ParserUseCase) Start() {
	uc.queue()
}

func (uc *ParserUseCase) CheckFeedURL(ctx context.Context, uri string) (string, error) {
	fp := uc.getParser()

	_, err := fp.ParseURLWithContext(uri, ctx)
	if errors.Is(err, gofeed.ErrFeedTypeNotDetected) {
		feedURL, err := uc.findFeedURL(ctx, uri)
		if err != nil {
			return "", err
		}
		return feedURL, nil
	} else if err != nil {
		if _, ok := err.(*url.Error); ok {
			return "", entity.ErrLoadURL
		}
		return "", err
	}

	return uri, nil
}

func (uc *ParserUseCase) LoadFeed(feed *entity.Feed) {
	uc.workersChan <- &feedJob{
		feed: feed,
	}
}

func (uc *ParserUseCase) findFeedURL(ctx context.Context, url string) (string, error) {
	data, err := uc.loadURL(ctx, url)
	if err != nil {
		return "", entity.ErrLoadURL
	}
	defer func() {
		_ = data.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return "", entity.ErrInternal
	}

	var (
		feedURL string
		exists  bool
	)

	res := doc.Find("link[type=\"application/rss+xml\"]")
	if feedURL, exists = res.Attr("href"); exists {
		return feedURL, nil
	}

	return "", nil
}

func (uc *ParserUseCase) loadURL(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, entity.ErrLoadURL
	}

	return resp.Body, nil
}

func (uc *ParserUseCase) init() {
	for i := 0; i < uc.cfg.WorkersPoolSize; i++ {
		go uc.worker()
	}
}

func (uc *ParserUseCase) queue() {
	go func() {
		uc.processFeeds()
		for {
			select {
			case <-uc.ctx.Done():
				return
			case <-time.After(time.Second * time.Duration(uc.cfg.FeedUpdateFrequency)):
				uc.processFeeds()
			}
		}
	}()
}

func (uc *ParserUseCase) processFeeds() {
	feeds, err := uc.feedsRepo.GetForUpdate(uc.ctx, uc.cfg.FeedUpdateFrequency)
	if err != nil {
		uc.log.Error().Err(err).Msg("failed to get feeds")
		return
	}

	for _, f := range feeds {
		uc.workersChan <- &feedJob{
			feed: f,
		}
	}
}

func (uc *ParserUseCase) worker() {
	fp := uc.getParser()

	for {
		select {
		case <-uc.ctx.Done():
			return
		case job := <-uc.workersChan:
			l := uc.log.Debug().Str("id", job.feed.ID).Str("site_url", job.feed.FeedURL)
			err := uc.jobProcess(job, fp)
			if err != nil {
				l.Err(err).Msg("processing job error")
				continue
			}
			l.Msg("processing job successful")
		}
	}
}

func (uc *ParserUseCase) jobProcess(job *feedJob, fp *gofeed.Parser) error {
	updFeed := &entity.Feed{
		ID:        job.feed.ID,
		UpdatedAt: structs.Ref(time.Now().UTC()),
	}

	feed, parseErr := fp.ParseURL(job.feed.FeedURL)
	if parseErr != nil {
		updFeed.ErrorCount++
	} else {
		updFeed.Title = feed.Title
	}

	items := make([]*entity.FeedItem, 0, len(feed.Items))
	var lastItemHash string

	for _, item := range feed.Items {
		hash := getItemHash(item)
		if hash == job.feed.LastItemHash {
			break
		}

		items = append(items, &entity.FeedItem{
			GUID:        item.GUID,
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
			Updated:     item.Updated,
			Published:   item.Published,
		})

		if len(lastItemHash) == 0 {
			lastItemHash = hash
		}
	}

	updFeed.LastItemHash = lastItemHash
	botUseCase.Send(updFeed, items)

	return nil
}

func (uc *ParserUseCase) getParser() *gofeed.Parser {
	fp := gofeed.NewParser()
	fp.UserAgent = userAgent
	return fp
}

func getItemHash(item *gofeed.Item) string {
	text := fmt.Sprintf("%s|%s|%s", item.Link, item.Updated, item.Published)
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
