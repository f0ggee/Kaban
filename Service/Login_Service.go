package Service

import (
	"Kaban/Dto"
	"Kaban/Service/Uttiltesss"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
)

func Login_Check(password string, hash_of_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash_of_password), []byte(password))
	if err != nil {
		slog.Error("Err", "err", err)
		return err

	}
	return nil

}
func Generate_Cookie(ctx context.Context, unic_Id int, dbe *sql.DB, r *http.Request, w http.ResponseWriter) error {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		slog.Error("Error generating id", err)
		return err

	}

	he := sha256.Sum256(b)

	h := hex.EncodeToString(he[:])

	_, err = dbe.ExecContext(ctx, "UPDATE cookies SET cookie = $1   WHERE  unic_id = $2", he, unic_Id)
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
	session.Values["cookie"] = h
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   100000,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if err := session.Save(r, w); err != nil {
		slog.Error("Cokie can't send", err)
		return err

	}

}

func Login_Service(s *Dto.Handler_Login, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx, canclel := Uttiltesss.Contexte()
	defer canclel()

	var (
		Unic_id  int
		password string
	)

	err := db.QueryRowContext(ctx, `SELECT unic_id ,password  FROM person WHERE email = $1`, s.Email).Scan(&Unic_id, &password)
	err = Uttiltesss.Err_Treate(err, w)
	if err != nil {
		slog.Error("Err", err)
		return
	}

	err = Login_Check(s.Password, password)
	if err != nil {
		slog.Error("Err", "err", err)
		return
	}
	err = Generate_Cookie(ctx, Unic_id, db, r, w)
	if err != nil {
		slog.Error("Err", "err", err)
		return
	}
	http.Redirect(w, r, "/main", http.StatusFound)

}
