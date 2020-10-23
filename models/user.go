package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	PasswordHash    string    `json:"-"`
	Created_at      time.Time `json:"-"`
	Updated_at      time.Time `json:"-"`
}
type UserDetail struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
	u.PasswordHash = string(pwdHash)

	now := time.Now()
	sql := "INSERT INTO users (name,email,password,created_at,updated_at) VALUES(?,?,?,?,?)"

	_, err = conn.Exec(sql, u.Name, u.Email, u.PasswordHash, now, now)
	return err

}

func (u *User) GetAuthToken() (string, error) {
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "secretfortoken")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	authToken, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return authToken, err
}
func (u *User) IsAuthenticated(conn *sql.DB) error {
	if strings.Trim(u.Email, " ") == "" || strings.Trim(u.Password, " ") == "" {
		return fmt.Errorf("Request tidak boleh kosong")
	}
	sql := "SELECT id,password FROM users WHERE email = ?"
	rows, err := conn.Query(sql, u.Email)
	if err != nil {
		return fmt.Errorf("Error Query Login")
	}
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.PasswordHash); err != nil {
			fmt.Println(err.Error())
			return fmt.Errorf("Invalid Login")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return fmt.Errorf("Invalid Email/Password")
	}

	return nil
}
func ExtracToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtracToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func GetPayload(r *http.Request) (uint64, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return 0, err
}

func (ud *UserDetail) GetUser(conn *sql.DB, r *http.Request) (*UserDetail, error) {
	userId, err := GetPayload(r)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	sql := "SELECT id,name,email FROM users WHERE id = ?"
	rows, err := conn.Query(sql, userId)
	if err != nil {
		return nil, fmt.Errorf("Error Query Data Users")
	}
	for rows.Next() {
		if err := rows.Scan(&ud.ID, &ud.Name, &ud.Email); err != nil {
			return nil, fmt.Errorf("Data gagal ditambahkan")
		}
	}
	return ud, nil
}
