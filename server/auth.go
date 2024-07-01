package main

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
    UserID int `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID int) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
