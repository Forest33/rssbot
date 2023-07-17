// Package database provides CRUD operations with database.
package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/database"
	"github.com/forest33/rssbot/pkg/logger"
	"github.com/forest33/rssbot/pkg/structs"
)

const (
	subscriptionsTable       = "subscriptions"
	subscriptionsTableFields = "id, user_id, feed_id, created_at"
)

// SubscriptionsRepository object capable of interacting with SubscriptionsRepository
type SubscriptionsRepository struct {
	db  *database.Database
	log *logger.Zerolog
}

// NewSubscriptionsRepository creates a new SubscriptionsRepository
func NewSubscriptionsRepository(db *database.Database, log *logger.Zerolog) *SubscriptionsRepository {
	return &SubscriptionsRepository{
		db:  db,
		log: log,
	}
}

type subscriptionDTO struct {
	ID        uuid.UUID `db:"id"`
	UserID    int64     `db:"user_id"`
	FeedID    uuid.UUID `db:"feed_id"`
	CreatedAt time.Time `db:"created_at"`
}

func newSubscriptionDTO(in *entity.Subscription) (dto *subscriptionDTO, err error) {
	dto = &subscriptionDTO{
		UserID: in.UserID,
	}

	dto.FeedID, err = uuid.Parse(in.FeedID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

func (dto *subscriptionDTO) entity() *entity.Subscription {
	return &entity.Subscription{
		ID:        dto.ID.String(),
		UserID:    dto.UserID,
		FeedID:    dto.FeedID.String(),
		CreatedAt: dto.CreatedAt,
	}
}

func identifierSubscriptionDTO(id string) (dto *subscriptionDTO, err error) {
	dto = &subscriptionDTO{}
	if id != "" {
		dto.ID, err = uuid.Parse(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return
}

// GetByID returns subscriptions by id
func (repo *SubscriptionsRepository) GetByID(ctx context.Context, id string) (*entity.Subscription, error) {
	dto, err := identifierSubscriptionDTO(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := repo.db.Connector.NamedQueryContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM %s
		WHERE id = :id`, subscriptionsTableFields, subscriptionsTable), dto)
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

	return nil, entity.ErrSubscriptionNotExists
}

// GetByUserID returns subscriptions by user id
func (repo *SubscriptionsRepository) GetByUserID(ctx context.Context, id int64) ([]*entity.Subscription, error) {
	var (
		query string
		args  []interface{}
		err   error
		dto   []*subscriptionDTO
	)

	attrs := []string{"user_id = :user_id"}
	mapper := map[string]interface{}{"user_id": id}

	query, args, err = repo.db.Connector.BindNamed(fmt.Sprintf(`
			SELECT %s 
			FROM %s
			WHERE %s
			ORDER BY created_at;`, subscriptionsTableFields, subscriptionsTable, strings.Join(attrs, ",")), mapper)
	if err != nil {
		return nil, err
	}
	if err := repo.db.Connector.SelectContext(ctx, &dto, query, args...); err != nil {
		return nil, err
	}

	return structs.Map(dto, func(s *subscriptionDTO) *entity.Subscription { return s.entity() }), nil
}

// GetByFeedID returns subscriptions by feed id
func (repo *SubscriptionsRepository) GetByFeedID(ctx context.Context, id string) ([]*entity.Subscription, error) {
	var (
		query string
		args  []interface{}
		err   error
		dto   []*subscriptionDTO
	)

	attrs := []string{"feed_id = :feed_id"}
	mapper := map[string]interface{}{"feed_id": id}

	query, args, err = repo.db.Connector.BindNamed(fmt.Sprintf(`
			SELECT %s 
			FROM %s
			WHERE %s
			ORDER BY created_at;`, subscriptionsTableFields, subscriptionsTable, strings.Join(attrs, ",")), mapper)
	if err != nil {
		return nil, err
	}
	if err := repo.db.Connector.SelectContext(ctx, &dto, query, args...); err != nil {
		return nil, err
	}

	return structs.Map(dto, func(s *subscriptionDTO) *entity.Subscription { return s.entity() }), nil
}

// Create creates new subscription
func (repo *SubscriptionsRepository) Create(ctx context.Context, in *entity.Subscription) (*entity.Subscription, error) {
	dto, err := newSubscriptionDTO(in)
	if err != nil {
		return nil, err
	}

	var (
		query string
		args  []interface{}
	)

	conn := repo.db.GetConnection(ctx)

	query, args, err = conn.BindNamed(fmt.Sprintf(`
			INSERT INTO %s (user_id, feed_id)
			VALUES (:user_id, :feed_id)
			RETURNING %s;`, subscriptionsTable, subscriptionsTableFields), dto)
	if err != nil {
		return nil, err
	}
	if err := conn.GetContext(ctx, dto, query, args...); err != nil {
		switch database.PgError(err) {
		case database.ErrDuplicateKey:
			return nil, entity.ErrSubscriptionAlreadyExists
		}
		return nil, err
	}

	return dto.entity(), nil
}

// Delete deletes user subscription
func (repo *SubscriptionsRepository) Delete(ctx context.Context, id string) error {
	dto, err := identifierSubscriptionDTO(id)
	if err != nil {
		return err
	}

	conn := repo.db.GetConnection(ctx)

	_, err = conn.NamedExecContext(ctx, fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = :id;`, subscriptionsTable), dto)
	if err != nil {
		return err
	}

	return nil
}
