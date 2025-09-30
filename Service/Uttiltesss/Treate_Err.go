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
		http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)
		slog.Info("func login1:timed out")
		return err

	case err == sql.ErrNoRows:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		slog.Info("func login2:no rows", err)
		return err

	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		slog.Info("func login3:no rows", err)
		return err

	default:
		slog.Info("func login4:ok")
		return nil

	}
}
