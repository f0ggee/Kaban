package Controller

import (
	"Kaban/Service"
	"github.com/gorilla/sessions"
	"log/slog"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("KEY"))

func Get_From(w http.ResponseWriter, r *http.Request) {

	slog.Info("eq")

	Service.F(w, r)

}
