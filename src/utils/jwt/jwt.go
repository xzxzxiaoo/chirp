package jwt

import (
	"chirp/src/errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	TokenExpired = errors.New("Token is expired")
	TokenInvalid = errors.New("Token is invalid")
)
var secret = []byte("myfateapisuhdfkej")

type Claims struct {
	Identity string `json:"identity"`
	jwt.StandardClaims
}

func GenerateToken(identity string) (string, error) {
	expires := time.Now().Add(time.Hour * 60).Unix()
	claims := Claims{
		identity,
		jwt.StandardClaims{
			ExpiresAt: expires,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenstr string) (*Claims, error) {
	if len(tokenstr) == 0 {
		return nil, TokenInvalid
	}

	token, err := jwt.ParseWithClaims(tokenstr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Printf("ParseToken Error: %s", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, TokenExpired
			}
		}
		return nil, TokenInvalid
	}

	return token.Claims.(*Claims), nil
}
