package JWT

import (
	"time"

	"github.com/Yscream/go-form-reg/pkg/DB"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int
	Name   string
	Lname  string
	jwt.StandardClaims
}

var (
	token_exp        = time.Now().Add(time.Minute * 15).Unix()
	HmacSampleSecret = []byte("secret")
)

func NewJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token_exp,
		},
		UserId: DB.GetId(email),
		Name:   DB.GetName(email),
		Lname:  DB.GetLname(email),
	})
	tokenString, err := token.SignedString(HmacSampleSecret)
	return tokenString, err
}
