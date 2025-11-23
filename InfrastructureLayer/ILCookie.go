package InfrastructureLayer

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Repositorys interface {
	UpDateCookie(ctx context.Context, Cookie string, UnicId int) (string, error)
}

type DB struct {
	Db *pgxpool.Pool
}

func (d *DB) UpDateCookie(ctx context.Context, Cookie string, UnicId int) (string, error) {
	var scrypt string

	_, err := d.Db.Exec(ctx, "UPDATE user_cookie SET cookie = $1 WHERE iduser = $2", Cookie, UnicId)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		slog.Error("Err don't find anathing", "err", err)
		return "", err

	case errors.Is(err, context.DeadlineExceeded):
		slog.Error("Context Dedline Exceed", "Error", err)
		return "", err

	}
	if err != nil {
		slog.Error("Err doesn't equal nil", "Err", err)
		return "", err
	}
	err = d.Db.QueryRow(ctx, "SELECT scrypt_salt  FROM person WHERE unic_id = $1", UnicId).Scan(&scrypt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		slog.Error("Doesn't cat select", "err", err)
		return "", err

	case errors.Is(err, context.DeadlineExceeded):
		slog.Error("Context Dedline Exceed", "Error", err)
		return "", err

	}
	if err != nil {
		slog.Error("Err was not equal err while choose", "Err", err)
		return "", err
	}

	return scrypt, err
}
