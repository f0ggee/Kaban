package Uttiltesss

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
)

func Err_Treate(err error, w http.ResponseWriter) error {
	switch {
	case err == context.DeadlineExceeded:
		slog.Info("func login1:timed out")
		return err

	case err == sql.ErrNoRows:
		slog.Info("func login2:no rows", err)
		return err

	case err != nil:
		slog.Info("func login3:no rows", err)
		return err

	default:
		slog.Info("func login4:ok")
		return nil

	}
}
