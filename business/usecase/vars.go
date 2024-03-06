package usecase

import (
	"context"
	"time"

	"github.com/forest33/rssbot/business/entity"
)

var (
	parserUseCase *ParserUseCase
	botUseCase    *BotUseCase
)

// FeedsRepo is the common interface implemented FeedsRepository methods
type FeedsRepo interface {
	Get(ctx context.Context, filter *entity.FeedFilter) ([]*entity.Feed, error)
	GetForUpdate(ctx context.Context, delay int64) ([]*entity.Feed, error)
	Create(ctx context.Context, in *entity.Feed) (*entity.Feed, error)
	Update(ctx context.Context, in *entity.Feed) (*entity.Feed, error)
	Delete(ctx context.Context, id string) error
}

// FeedItemsRepo is the common interface implemented FeedItemsRepository methods
type FeedItemsRepo interface {
	Create(ctx context.Context, in *entity.FeedItem) (*entity.FeedItem, error)
	Delete(ctx context.Context, feedID string, toDate time.Time) error
}

func SetParserUseCase(uc *ParserUseCase) {
	parserUseCase = uc
}

func SetBotUseCase(uc *BotUseCase) {
	botUseCase = uc
}
