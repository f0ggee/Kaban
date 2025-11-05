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
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"time"
)

func Login_Check(password string, hash_of_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash_of_password), []byte(password))
	if err != nil {
		slog.Error("Err", "err", err)
		return err

	}
	return nil

}
func Generate_Cookie(ctx context.Context, db *pgxpool.Pool, unic_Id int, r *http.Request, w http.ResponseWriter) error {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		slog.Error("Error generating id", err)
		return err

	}
	he := hex.EncodeToString(b)

	_, err = db.Exec(context.Background(), "UPDATE cookies SET cookie = $1   WHERE  unic_id = $2", he, unic_Id)
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

func Login_Service(s *Dto.Handler_Login, w http.ResponseWriter, r *http.Request) {
	db, err := Connect_to_BD.Connect()
	if err != nil {
		slog.Error("Err for connect to bd in login", err)
		return
	}
	ctx, canclel := Uttiltesss.Contexte()
	defer canclel()

	var (
		Unic_id  int
		password string
	)

	err = db.QueryRow(context.Background(), `SELECT unic_id ,password, FROM person WHERE email=$1`, s.Email).Scan(&Unic_id, &password)
	err = Uttiltesss.Err_Treate(err, w)
	if err != nil {
		slog.Error("Func login 1 ", err)
		return
	}

	err = Login_Check(s.Password, password)
	if err != nil {
		slog.Error("func login 2", "err", err)
		return
	}
	err = Generate_Cookie(ctx, db, Unic_id, r, w)
	if err != nil {
		slog.Error("func login 3", "err", err)
		return
	}
	http.Redirect(w, r, "/main", http.StatusFound)

}
