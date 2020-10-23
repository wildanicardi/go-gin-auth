package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	PasswordHash    string    `json:"password_hash"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}

func (u *User) Register(conn *sql.DB) error {
	if len(u.Password) < 6 || len(u.PasswordConfirm) < 6 {
		return fmt.Errorf("Password Terlalu Pendek")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Password Tidak cocok")
	}


	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Error Hash Akun")
	}
	u.Password = string(pwdHash)

	now := time.Now()
	sql := "INSERT INTO users (name,email,password,created_at,updated_at) VALUES(?,?,?,?,?)"

	_, err = conn.Exec(sql, u.Name, u.Email, u.Password, now, now)
	return err

}

func (u *User) GetAuthToken() (string, error) {
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "secretfortoken")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	authToken, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return authToken, err
}
