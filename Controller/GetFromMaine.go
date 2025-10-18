package Controller

import (
	"github.com/gorilla/sessions"
	"log/slog"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("KEY"))

func Get_From(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "cant' treate", http.StatusNotFound)
		slog.Info("notfoud")
		return
	}

	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		http.Error(w, "cookie dont sen", http.StatusUnauthorized)
	}
	slog.Info("eae")
	if session.Options.MaxAge == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	sa, ok := session.Values["cookie"]
	if ok {
		slog.Info("dsa", sa)
		http.Redirect(w, r, "/main", http.StatusMovedPermanently)

	}

}
