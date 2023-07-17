// Package database provides CRUD operations with database.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/database"
	"github.com/forest33/rssbot/pkg/database/types"
	"github.com/forest33/rssbot/pkg/logger"
	"github.com/forest33/rssbot/pkg/structs"
)

const (
	feedsTable       = "feeds"
	feedsTableFields = "id, title, site_url, feed_url, last_item_hash, error_count, created_at, updated_at"
)

// FeedsRepository object capable of interacting with FeedsRepository
type FeedsRepository struct {
	db  *database.Database
	log *logger.Zerolog
}

// NewFeedsRepository creates a new FeedsRepository
func NewFeedsRepository(db *database.Database, log *logger.Zerolog) *FeedsRepository {
	return &FeedsRepository{
		db:  db,
		log: log,
	}
}

type feedDTO struct {
	ID           uuid.UUID    `db:"id"`
	Title        string       `db:"title"`
	SiteURL      string       `db:"site_url"`
	FeedURL      string       `db:"feed_url"`
	ErrorCount   int          `db:"error_count"`
	LastItemHash string       `db:"last_item_hash"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}

func newFeedDTO(in *entity.Feed) (dto *feedDTO, err error) {
	dto = &feedDTO{
		Title:   in.Title,
		SiteURL: in.SiteURL,
		FeedURL: in.FeedURL,
	}
	return
}

func (dto *feedDTO) entity() *entity.Feed {
	return &entity.Feed{
		ID:           dto.ID.String(),
		Title:        dto.Title,
		SiteURL:      dto.SiteURL,
		FeedURL:      dto.FeedURL,
		LastItemHash: dto.LastItemHash,
		ErrorCount:   dto.ErrorCount,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    types.SQLToRefTime(dto.UpdatedAt),
	}
}

func identifierFeedDTO(id string) (dto *feedDTO, err error) {
	dto = &feedDTO{}
	if id != "" {
		dto.ID, err = uuid.Parse(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return
}

// GetByID returns feed by id
func (repo *FeedsRepository) GetByID(ctx context.Context, id string) (*entity.Feed, error) {
	dto, err := identifierFeedDTO(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := repo.db.Connector.NamedQueryContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM %s
		WHERE id = :id`, feedsTableFields, feedsTable), dto)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	if rows.Next() {
		err := rows.StructScan(&dto)
		if err != nil {
			return nil, err
		}
		return dto.entity(), nil
	}

	return nil, entity.ErrFeedNotExists
}

// Get returns feeds
func (repo *FeedsRepository) Get(ctx context.Context, filter *entity.FeedFilter) ([]*entity.Feed, error) {
	var (
		query  string
		args   []interface{}
		err    error
		dto    []*feedDTO
		attrs  = make([]string, 0, 1)
		mapper = make(map[string]interface{}, 1)
	)

	if filter.ID != nil {
		attrs = append(attrs, "id IN(:id)")
		mapper["id"] = filter.ID
	}

	query, args, err = sqlx.Named(fmt.Sprintf(`
			SELECT %s 
			FROM %s
			WHERE %s;`, feedsTableFields, feedsTable, strings.Join(attrs, ",")), mapper)
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

	return structs.Map(dto, func(w *feedDTO) *entity.Feed { return w.entity() }), nil
}

// GetForUpdate returns feeds that need to be updated
func (repo *FeedsRepository) GetForUpdate(ctx context.Context, delay int64) ([]*entity.Feed, error) {
	var (
		query string
		args  []interface{}
		err   error
		dto   []*feedDTO
	)

	attrs := []string{"updated_at <= :updated_at"}
	mapper := map[string]interface{}{"updated_at": time.Unix(time.Now().UTC().Unix()-delay, 0)}

	query, args, err = repo.db.Connector.BindNamed(fmt.Sprintf(`
			SELECT %s 
			FROM %s
			WHERE updated_at IS NULL OR %s
			ORDER BY updated_at DESC;`, feedsTableFields, feedsTable, strings.Join(attrs, ",")), mapper)
	if err != nil {
		return nil, err
	}
	if err := repo.db.Connector.SelectContext(ctx, &dto, query, args...); err != nil {
		return nil, err
	}

	return structs.Map(dto, func(w *feedDTO) *entity.Feed { return w.entity() }), nil
}

// Create creates new feed
func (repo *FeedsRepository) Create(ctx context.Context, in *entity.Feed) (*entity.Feed, error) {
	dto, err := newFeedDTO(in)
	if err != nil {
		return nil, err
	}

	var (
		query string
		args  []interface{}
	)

	conn := repo.db.GetConnection(ctx)

	query, args, err = conn.BindNamed(fmt.Sprintf(`
			INSERT INTO %s (title, site_url, feed_url)
			VALUES (:title, :site_url, :feed_url)
			ON CONFLICT (feed_url) DO UPDATE SET updated_at	= NOW()
			RETURNING %s;`, feedsTable, feedsTableFields), dto)
	if err != nil {
		return nil, err
	}
	if err := conn.GetContext(ctx, dto, query, args...); err != nil {
		return nil, err
	}

	return dto.entity(), nil
}

// Update updates feed item
func (repo *FeedsRepository) Update(ctx context.Context, in *entity.Feed) (*entity.Feed, error) {
	dto, err := identifierFeedDTO(in.ID)
	if err != nil {
		return nil, err
	}

	attrs := make([]string, 0, 4)
	mapper := make(map[string]interface{}, 4)

	if in.Title != "" {
		attrs = append(attrs, "title = :title")
		mapper["title"] = in.Title
	}
	if in.LastItemHash != "" {
		attrs = append(attrs, "last_item_hash = :last_item_hash")
		mapper["last_item_hash"] = in.LastItemHash
	}
	if in.ErrorCount >= 0 {
		attrs = append(attrs, "error_count = :error_count")
		mapper["error_count"] = in.ErrorCount
	}
	if in.UpdatedAt != nil {
		attrs = append(attrs, "updated_at = :updated_at")
		mapper["updated_at"] = in.UpdatedAt
	}
	if len(attrs) == 0 {
		return repo.GetByID(ctx, in.ID)
	}

	attrs = append(attrs, "id = :id")
	mapper["id"] = dto.ID

	var (
		query string
		args  []interface{}
	)

	conn := repo.db.GetConnection(ctx)

	query, args, err = conn.BindNamed(fmt.Sprintf(`
			UPDATE %s SET %s
			WHERE id = :id
			RETURNING %s;`, feedsTable, strings.Join(attrs, ","), feedsTableFields), mapper)
	if err != nil {
		return nil, err
	}
	if err := conn.GetContext(ctx, dto, query, args...); err != nil {
		return nil, err
	}

	return dto.entity(), nil
}

// Delete deletes feed item
func (repo *FeedsRepository) Delete(ctx context.Context, id string) error {
	dto, err := identifierFeedDTO(id)
	if err != nil {
		return err
	}

	conn := repo.db.GetConnection(ctx)

	_, err = conn.NamedExecContext(ctx, fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = :id;`, feedsTable), dto)
	if err != nil {
		return err
	}

	return nil
}
