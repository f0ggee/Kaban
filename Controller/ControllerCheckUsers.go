package Controller

import (
	"github.com/gorilla/sessions"
	"log/slog"
	"net"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("KEY"))

func Get_From(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "cant' treate", http.StatusNotFound)
		slog.Info("notfoud")
		return
	}

	clientIP := r.RemoteAddr

	host, _, err := net.SplitHostPort(clientIP)
	if err != nil {
		slog.Info("ee")
		return
	}
	clientIP = host

	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		http.Error(w, "cookie dont sen", http.StatusUnauthorized)
	}
	if session.Options.MaxAge == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	sa, ok := session.Values["cookie"]
	if ok {

		slog.Info("Cookie of user", sa)
		http.Redirect(w, r, "/main", http.StatusMovedPermanently)

	}

}
