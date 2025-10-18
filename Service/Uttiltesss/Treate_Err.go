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
		slog.Info("func login1:timed out", err)
		return err

	case err == sql.ErrNoRows:
		slog.Info("func login2:no rows", err)
		return err

	}

	return nil

}
