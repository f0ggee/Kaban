package Controller

import (
	"Kaban/Dto"
	"Kaban/Service/Handlers"
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

	Handlers.Login_Service(sa, w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(w)

}
