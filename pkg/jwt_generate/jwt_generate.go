package jwt_generate

import (
	"time"

	"github.com/Yscream/go-form-reg/pkg/store"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int
	User   string
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
		UserId: store.GetId(email),
		User:   store.GetUser(email),
	})
	tokenString, err := token.SignedString(HmacSampleSecret)
	return tokenString, err
}
