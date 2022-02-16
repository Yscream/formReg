package JWT

import (
	"fmt"
	"log"

	"github.com/Yscream/go-form-reg/pkg/DB"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/golang-jwt/jwt"
)

func SaveToken(user *models.LoginUser) {
	db := DB.GetConnection()

	defer db.Close()

	token, err := NewJWT(user.Email)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("INSERT INTO tokens (users_id, token)  VALUES($1, $2)", DB.GetId(user.Email), token)
}

func DeleteToken(token string) {
	db := DB.GetConnection()

	defer db.Close()

	db.Exec("DELETE FROM tokens WHERE token=$1", token)
}

func SendToken(user *models.LoginUser) string {
	db := DB.GetConnection()

	defer db.Close()

	var token string
	err := db.Get(&token, "SELECT token FROM tokens WHERE users_id=$1", DB.GetId(user.Email))
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func ParseJWT(tokenStr string, hmacSecret []byte) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("ololo")
		return nil
	} else {
		fmt.Println(err)
		return err
	}
}
