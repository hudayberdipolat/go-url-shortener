package jwtToken

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

var SecretKey = []byte("das#jd!ahDjSwr$we$ry$wbw_we^t*&^$%^#$sa)adEd&$sda")

func GenerateToken(username string, userID int) (string, error) {
	claims := Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
