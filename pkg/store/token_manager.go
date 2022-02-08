package store

import (
	"fmt"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/Yscream/go-form-reg/jwt_generate"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

func SaveToken(user *models.LoginUser) {
	conn := configs.GetConfig()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	token, err := jwt_generate.NewJWT(user.Email)

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
	conn := configs.GetConfig()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var token string
	err = db.QueryRow("SELECT token FROM tokens WHERE users_id=$1", GetId(user.Email)).Scan(&token)
	if err != nil {
		fmt.Println("lul, wtf are you doing?")
	}
	return token
}

func ParseJWT(tokenStr string, hmacSampleSecret []byte) error {
	fmt.Println("sdgsd")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("ololo")
		return nil
	} else {
		fmt.Println("error")
		fmt.Println(err)
		return err
	}
}
