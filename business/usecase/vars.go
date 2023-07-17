package usecase

import (
	"context"

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

func SetParserUseCase(uc *ParserUseCase) {
	parserUseCase = uc
}

func SetBotUseCase(uc *BotUseCase) {
	botUseCase = uc
}
