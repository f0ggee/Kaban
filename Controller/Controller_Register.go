package Controller

import (
	"Kaban/Dto"
	"Kaban/Service/Handlers"
	"encoding/json"
	"log/slog"
	"net/http"
)

func chehkjson_Register(r *http.Request) (*Dto.Handler_Registerr, error) {
	var err error
	var e Dto.Handler_Registerr

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &e, err
}

func Controller_Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		slog.Error("Error from Controller_register, method don't allow ", "err")
		http.Error(w, "M don't allow", http.StatusNotFound)
		return
	}

	t, err := chehkjson_Register(r)
	if err != nil {
		slog.Error("Error from Controller_register", "err", err)
		return
	}

	Handlers.Register_Service(t, w, r)
}
