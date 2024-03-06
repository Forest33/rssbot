// Package database provides CRUD operations with database.
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/database"
	"github.com/forest33/rssbot/pkg/logger"
	"github.com/forest33/rssbot/pkg/structs"
)

const (
	feedItemsTable       = "feed_items"
	feedItemsTableFields = "id, feed_id, item_hash, created_at"
)

// FeedItemsRepository object capable of interacting with FeedItemsRepository
type FeedItemsRepository struct {
	db  *database.Database
	log *logger.Zerolog
}

// NewFeedItemsRepository creates a new FeedItemsRepository
func NewFeedItemsRepository(db *database.Database, log *logger.Zerolog) *FeedItemsRepository {
	return &FeedItemsRepository{
		db:  db,
		log: log,
	}
}

type feedItemDTO struct {
	ID        uuid.UUID `db:"id"`
	FeedID    uuid.UUID `db:"feed_id"`
	ItemHash  string    `db:"item_hash"`
	CreatedAt time.Time `db:"created_at"`
}

func newFeedItemDTO(in *entity.FeedItem) (dto *feedItemDTO, err error) {
	dto = &feedItemDTO{
		ItemHash: in.ItemHash,
	}

	dto.FeedID, err = uuid.Parse(in.FeedID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

func (dto *feedItemDTO) entity() *entity.FeedItem {
	return &entity.FeedItem{
		ID:        dto.ID.String(),
		FeedID:    dto.FeedID.String(),
		ItemHash:  dto.ItemHash,
		CreatedAt: dto.CreatedAt,
	}
}

func identifierFeedItemDTO(feedID string) (dto *feedItemDTO, err error) {
	dto = &feedItemDTO{}
	dto.FeedID, err = uuid.Parse(feedID)
	return
}

// GetByFeedID returns feed items by feed_id
func (repo *FeedItemsRepository) GetByFeedID(ctx context.Context, id string) ([]*entity.FeedItem, error) {
	var (
		query string
		args  []interface{}
		err   error
		dto   []*feedItemDTO
	)

	query, args, err = sqlx.Named(fmt.Sprintf(`
			SELECT %s 
			FROM %s
			WHERE feed_id = :id;`, feedItemsTableFields, feedItemsTable), map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = repo.db.Connector.Rebind(query)
	if err := repo.db.Connector.SelectContext(ctx, &dto, query, args...); err != nil {
		return nil, err
	}

	return structs.Map(dto, func(w *feedItemDTO) *entity.FeedItem { return w.entity() }), nil
}

// Create creates new feed item
func (repo *FeedItemsRepository) Create(ctx context.Context, in *entity.FeedItem) (*entity.FeedItem, error) {
	dto, err := newFeedItemDTO(in)
	if err != nil {
		return nil, err
	}

	var (
		query string
		args  []interface{}
	)

	conn := repo.db.GetConnection(ctx)

	query, args, err = conn.BindNamed(fmt.Sprintf(`
			INSERT INTO %s (feed_id, item_hash)
			VALUES (:feed_id, :item_hash)
			RETURNING %s;`, feedItemsTable, feedItemsTableFields), dto)
	if err != nil {
		return nil, err
	}
	if err := conn.GetContext(ctx, dto, query, args...); err != nil {
		var e *pq.Error
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return nil, nil
		}
		return nil, err
	}

	return dto.entity(), nil
}

// Delete deletes feed item
func (repo *FeedItemsRepository) Delete(ctx context.Context, feedID string, toDate time.Time) error {
	dto, err := identifierFeedItemDTO(feedID)
	if err != nil {
		return err
	}
	dto.CreatedAt = toDate

	conn := repo.db.GetConnection(ctx)

	_, err = conn.NamedExecContext(ctx, fmt.Sprintf(`
			DELETE FROM %s
			WHERE feed_id = :feed_id AND created_at < :created_at;`, feedItemsTable), dto)
	if err != nil {
		return err
	}

	return nil
}
