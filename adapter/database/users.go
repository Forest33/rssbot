// Package database provides CRUD operations with database.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/database"
	"github.com/forest33/rssbot/pkg/database/types"
	"github.com/forest33/rssbot/pkg/logger"
	"github.com/forest33/rssbot/pkg/structs"
)

const (
	usersTable       = "users"
	usersTableFields = "id, chat_id, first_name, last_name, username, is_bot, language, created_at, updated_at"
)

// UsersRepository object capable of interacting with UsersRepository
type UsersRepository struct {
	db  *database.Database
	log *logger.Zerolog
}

// NewUsersRepository creates a new UsersRepository
func NewUsersRepository(db *database.Database, log *logger.Zerolog) *UsersRepository {
	return &UsersRepository{
		db:  db,
		log: log,
	}
}

type userDTO struct {
	ID        int64        `db:"id"`
	ChatID    int64        `db:"chat_id"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	UserName  string       `db:"username"`
	IsBot     bool         `db:"is_bot"`
	Language  string       `db:"language"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func newUserDTO(in *entity.User) (dto *userDTO, err error) {
	dto = &userDTO{
		ID:        in.ID,
		ChatID:    in.ChatID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		UserName:  in.UserName,
		IsBot:     in.IsBot,
		Language:  in.Language,
	}
	return
}

func (dto *userDTO) entity() *entity.User {
	return &entity.User{
		ID:        dto.ID,
		ChatID:    dto.ChatID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		UserName:  dto.UserName,
		IsBot:     dto.IsBot,
		Language:  dto.Language,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: types.SQLToRefTime(dto.UpdatedAt),
	}
}

// GetByID returns user by id
func (repo *UsersRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	dto := &userDTO{ID: id}

	rows, err := repo.db.Connector.NamedQueryContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM %s
		WHERE id = :id`, usersTableFields, usersTable), dto)
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

	return nil, entity.ErrUserNotExists
}

// Get returns users
func (repo *UsersRepository) Get(ctx context.Context, filter *entity.UsersFilter) ([]*entity.User, error) {
	var (
		query  string
		args   []interface{}
		err    error
		dto    []*userDTO
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
			WHERE %s;`, usersTableFields, usersTable, strings.Join(attrs, ",")), mapper)
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

	return structs.Map(dto, func(u *userDTO) *entity.User { return u.entity() }), nil
}

// Create creates new user
func (repo *UsersRepository) Create(ctx context.Context, in *entity.User) (*entity.User, error) {
	dto, err := newUserDTO(in)
	if err != nil {
		return nil, err
	}

	var (
		query string
		args  []interface{}
	)

	conn := repo.db.GetConnection(ctx)

	query, args, err = conn.BindNamed(fmt.Sprintf(`
			INSERT INTO %s (id, chat_id, first_name, last_name, username, is_bot, language)
			VALUES (:id, :chat_id, :first_name, :last_name, :username, :is_bot, :language)
			ON CONFLICT (id) DO UPDATE SET
				updated_at 	= NOW(),
			    chat_id		= :chat_id,
			    first_name 	= :first_name,
			    last_name 	= :last_name,
			    username 	= :username,
			    is_bot 		= :is_bot,
			    language 	= :language
			RETURNING %s;`, usersTable, usersTableFields), dto)
	if err != nil {
		return nil, err
	}
	if err := conn.GetContext(ctx, dto, query, args...); err != nil {
		return nil, err
	}

	return dto.entity(), nil
}
