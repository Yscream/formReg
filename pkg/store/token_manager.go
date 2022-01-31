package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Yscream/formReg/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserId int
}

var (
	token_exp        = int64(time.Hour * 12)
	hmacSampleSecret []byte
)

func NewJWTtoken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token_exp,
		},
		UserId: GetId(email),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
	return tokenString, err
}

func SaveToken(user *models.LoginUser) {
	conn := GetConfig()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	token, err := NewJWTtoken(user.Email)

	if err != nil {
		fmt.Println("lul, wtf are you doing?")
	}

	tokens, err := db.Query("INSERT INTO tokens (users_id, token)  VALUES($1, $2)", GetId(user.Email), token)
	if err != nil {
		panic(err)
	}

	defer tokens.Close()
}

func SendToken(user *models.LoginUser) string {
	conn := GetConfig()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var token string
	err = db.QueryRow("SELECT token FROM tokens WHERE users_id=$1", GetId(user.Email)).Scan(&token)
	if err != nil {
		fmt.Println(token)
	}
	fmt.Println("from 90", token)
	return token
}

func ParseJWTtoken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte("sdsad"), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
