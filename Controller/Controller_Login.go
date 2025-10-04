package Controller

import (
	"Kaban/Dto"
	"Kaban/Service/Handlers"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
)

func chehkjson(r *http.Request) (*Dto.Handler_Login, error) {
	var err error
	var e Dto.Handler_Login

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &e, err
}

func Loging(w http.ResponseWriter, r *http.Request) {

	var db *sql.DB
	if r.Method != http.MethodPost {
		http.Error(w, "Erro", http.StatusUnauthorized)
		slog.Error("Error", "err")
		return
	}

	sa, err := chehkjson(r)
	if err != nil {
		slog.Error("Err", "error", err)
		return

	}

	Handlers.Login_Service(sa, db, w, r)

}
