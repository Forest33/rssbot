package usecase

import (
	"context"
	"time"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/logger"
)

// CleanerUseCase object capable of interacting with CleanerUseCase
type CleanerUseCase struct {
	ctx       context.Context
	cfg       *entity.CleanerConfig
	log       *logger.Zerolog
	feedsRepo FeedsRepo
	itemsRepo FeedItemsRepo
}

// NewCleanerUseCase creates a new CleanerUseCase
func NewCleanerUseCase(ctx context.Context, cfg *entity.CleanerConfig, log *logger.Zerolog, feedsRepo FeedsRepo, itemsRepo FeedItemsRepo) (*CleanerUseCase, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	uc := &CleanerUseCase{
		ctx:       ctx,
		cfg:       cfg,
		log:       log,
		feedsRepo: feedsRepo,
		itemsRepo: itemsRepo,
	}

	return uc, nil
}

func (uc *CleanerUseCase) Start() {
	ticker := time.NewTicker(time.Second * time.Duration(uc.cfg.Interval))
	go func() {
		select {
		case <-ticker.C:
			uc.clean()
		case <-uc.ctx.Done():
			return
		}
	}()
}

func (uc *CleanerUseCase) clean() {
	feeds, err := uc.feedsRepo.Get(uc.ctx, &entity.FeedFilter{})
	if err != nil {
		uc.log.Error().Err(err).Msg("failed to get feeds")
		return
	}
	for i := range feeds {
		if feeds[i].UpdatedAt == nil {
			continue
		}
		toDate := feeds[i].UpdatedAt.AddDate(0, 0, -uc.cfg.FeedItemTTL)
		if err := uc.itemsRepo.Delete(uc.ctx, feeds[i].ID, toDate); err != nil {
			uc.log.Error().Err(err).Msg("failed to clean feed items")
		}
	}
}
