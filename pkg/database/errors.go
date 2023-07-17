package database

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	pgErrorCodeDuplicateKey = "23505"
)

var (
	ErrDuplicateKey = errors.New("duplicate key value violates unique constraint")
)

func PgError(err error) error {
	if uErr := errors.Unwrap(err); uErr != nil {
		err = uErr
	}
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case pgErrorCodeDuplicateKey:
			return ErrDuplicateKey
		}
	}
	return err
}
