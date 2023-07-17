// Package database provides low level operations with database.
package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/pkg/logger"
)

// Database object capable of interacting with Database
type Database struct {
	Connector     *sqlx.DB
	cfg           *entity.DatabaseConfig
	binDataConfig *BinDataConfig
	log           *logger.Zerolog
}

type txContextKey struct{}

type Connection interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// NewConnector creates a new Database
func NewConnector(cfg *entity.DatabaseConfig, binDataConfig *BinDataConfig, log *logger.Zerolog) (*Database, error) {
	db := &Database{
		cfg:           cfg,
		log:           log,
		binDataConfig: binDataConfig,
	}

	log.Debug().
		Str("driver", cfg.Driver).
		Str("host", cfg.Host).
		Int("port", cfg.Port).
		Str("username", cfg.Username).
		Str("database", cfg.Database).
		Msg("initialize database")

	var err error
	db.Connector, err = sqlx.Connect(cfg.Driver, db.getDSN(cfg))
	if err != nil {
		return nil, err
	}

	if _, err := db.migrate(); err != nil {
		return nil, err
	}

	return db, nil
}

// Close closes database connection
func (db *Database) Close() {
	if err := db.Connector.Close(); err != nil {
		db.log.Error().Msgf("failed to close database: %v", err)
	}
}

// BeginTransaction begins a transaction
func (db *Database) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := db.Connector.Beginx()
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, txContextKey{}, tx), nil
}

// CommitTransaction commits the transaction
func (db *Database) CommitTransaction(ctx context.Context, txErr error) {
	tx, ok := ctx.Value(txContextKey{}).(*sqlx.Tx)
	if !ok {
		return
	}

	if txErr != nil {
		if err := tx.Rollback(); err != nil {
			db.log.Error().Err(err).Msg("failed to rollback transaction")
			return
		}
	}

	if err := tx.Commit(); err != nil {
		db.log.Error().Err(err).Msg("failed to commit transaction")
	}
}

// AbortTransaction aborts the transaction
func (db *Database) AbortTransaction(ctx context.Context) error {
	tx, ok := ctx.Value(txContextKey{}).(*sqlx.Tx)
	if !ok {
		return nil
	}

	if err := tx.Rollback(); err != nil {
		db.log.Error().Err(err).Msg("failed to rollback transaction")
		return err
	}

	return nil
}

func (db *Database) GetConnection(ctx context.Context) Connection {
	tx, ok := ctx.Value(txContextKey{}).(*sqlx.Tx)
	if !ok {
		return db.Connector
	}
	return tx
}

func (db *Database) getDSN(cfg *entity.DatabaseConfig) string {
	switch cfg.Driver {
	case SQLiteDriver:
		return cfg.Database
	case PostgresDriver:
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode)
	default:
		db.log.Fatal(fmt.Sprintf("unknown database driver %s", cfg.Driver))
	}
	return ""
}
