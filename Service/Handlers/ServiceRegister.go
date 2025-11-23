package Handlers

import (
	"Kaban/Dto"
	"Kaban/Service/Connect_to_BD"
	"Kaban/Service/Uttiltesss"
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/scrypt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func aenerate_Cookie(ctx context.Context, unic_Id int, dbe *pgxpool.Pool, r *http.Request, w http.ResponseWriter) error {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		slog.Error("Error generating id", err)
		return err

	}
	he := hex.EncodeToString(b)

	err = dbe.QueryRow(context.Background(), "INSERT INTO user_cookie(iduser,cookie) VALUES($1,$2)", unic_Id, he).Scan(nil)
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
	session.Values["SW"] = he
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

func Register_Service(de *Dto.Handler_Registerr, w http.ResponseWriter, r *http.Request) {

	ctx, Cancel := Uttiltesss.Contexte()
	defer Cancel()

	db, err := Connect_to_BD.Connect()
	if err != nil {
		slog.Error("Err_from_register 1 ", err)
		return
	}

	if len(de.Password) < 3 {
		http.Error(w, "Password are small", http.StatusUnauthorized)
		slog.Error("Err password are smal")
		return
	}
	slog.Info(de.Email)
	if !strings.Contains(de.Email, "@") {
		http.Error(w, "Person name must contain @", http.StatusBadRequest)
		slog.Info("Func register2")
		return
	}

	var existingPerson bool

	err = db.QueryRow(context.Background(), "SELECT EXISTS (select 1 FROM person WHERE email=$1)", de.Email, de.Name).Scan(&existingPerson)

	err = Uttiltesss.Err_Treate(err, w)
	if err != nil {
		http.Error(w, "Can't treat request", http.StatusNotFound)
		slog.Error("Error from register 2", err)
		return
	}

	if existingPerson {
		http.Error(w, "person already exists", http.StatusConflict)
		slog.Info("Func register5:", err)
		return
	}

	te, err := Uttiltesss.HashPassowrd(de.Password)
	if err != nil {
		slog.Error("Err from register 3", err)
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

	err = db.QueryRow(context.Background(), "INSERT INTO person(name,email,password,scrypt_salt,created_at) VALUES ($1,$2,$3,$4,$5) RETURNING unic_id", de.Name, de.Email, te, D4, time.Now()).Scan(&unic_id)

	if err != nil {
		http.Error(w, "Person already exist", http.StatusUnauthorized)
		slog.Error("Err from register 4", err)
		return
	}

	err = aenerate_Cookie(ctx, unic_id, db, r, w)
	if err != nil {
		slog.Error("Err from register 5", err)
		http.Error(w, "Err", http.StatusNotFound)
		return
	}
	time.Sleep(2 * time.Second)

	http.Redirect(w, r, "/main", http.StatusFound)

}
