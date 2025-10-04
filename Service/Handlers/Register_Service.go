package Handlers

import (
	"Kaban/Dto"
	"Kaban/Service/Uttiltesss"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/scrypt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func aenerate_Cookie(ctx context.Context, unic_Id int, dbe *sql.DB, r *http.Request, w http.ResponseWriter) error {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		slog.Error("Error generating id", err)
		return err

	}
	he := hex.EncodeToString(b)

	_, err = dbe.ExecContext(ctx, "UPDATE person SET cookie = $1   WHERE  unic_id = $2", he, unic_Id)
	err = Uttiltesss.Err_Treate(err, w)
	if err != nil {
		slog.Error("Err", err)
		return err
	}

	session, err := store.Get(r, "token1")
	if err != nil {
		slog.Error("cookie don't send", err)
		http.Error(w, "cookie dont sen", http.StatusUnauthorized)
		return err
	}
	session.Values["cookie"] = he
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int((24 * time.Hour).Seconds()),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if err := session.Save(r, w); err != nil {
		slog.Error("Cokie can't send", err)
		return err

	}
	return nil

}

func Register_Service(de *Dto.Handler_Registerr, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx, Cancel := Uttiltesss.Contexte()
	defer Cancel()

	var err error

	if len(de.Password) < 3 {
		http.Error(w, "Password are small", http.StatusUnauthorized)
		slog.Error("Err password are smal")
		return
	}
	if !strings.Contains(de.Email, "@") {
		http.Error(w, "Person name must contain @", http.StatusBadRequest)
		slog.Info("Func register2:", err)
		return
	}

	var existingPerson bool

	err = db.QueryRowContext(ctx, "SELECT EXISTS (select 1 FROM person WHERE email=$1", de.Email, de.Name).Scan(&existingPerson)

	err = Uttiltesss.Err_Treate(err, w)
	if err != nil {
		http.Error(w, "Can't treat request", http.StatusNotFound)
		slog.Error("Error", err)
		return
	}

	if existingPerson {
		http.Error(w, "person already exists", http.StatusConflict)
		slog.Info("Func register5:", err)
		return
	}

	te, err := Uttiltesss.HashPassowrd(de.Password)
	if err != nil {
		slog.Error("Err", err)
		return
	}

	salt := make([]byte, 16)
	rand.Read(salt)
	D3, err := scrypt.Key([]byte(de.Password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		slog.Error("Error cant create password", err)
		return
	}
	D4 := hex.EncodeToString(D3)

	var unic_id int

	err = db.QueryRowContext(ctx, "INSERT INTO person(name,email,password,scrypt_salt) VALUES ($1,$2,$3,$4) RETURNING unic_id", de.Name, de.Email, te, D4).Scan(&unic_id)

	if err != nil {
		slog.Error("Err", err)
		return
	}

	err = aenerate_Cookie(ctx, unic_id, db, r, w)
	if err != nil {
		slog.Error("Err", err)
		http.Error(w, "Err", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/main", http.StatusFound)

}
