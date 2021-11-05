package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthClaim struct {
	UserId uint64 `json:"userId"`
	jwt.StandardClaims
}

var secret = []byte("It is my secret")

const TokenExpireDuration = 2 * time.Hour

func GenToken(userId uint64) (string, error) {
	c := AuthClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "TEST001",
		},
	}
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (*AuthClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaim{}, func(tk *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("Invalid token ")
}
