package JWT

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int
	Name   string
	Lname  string
	jwt.StandardClaims
}

var (
	token_exp        = time.Now().Add(time.Minute * 5).Unix()
	HmacSampleSecret = []byte("secret")
)

func NewJWT(email, name, lname string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token_exp,
		},
		UserId: id,
		Name:   name,
		Lname:  lname,
	})
	tokenString, err := token.SignedString(HmacSampleSecret)
	return tokenString, err
}
