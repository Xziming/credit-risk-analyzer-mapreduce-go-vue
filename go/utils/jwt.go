package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
    expirationTime := time.Now().Add(72 * time.Hour)

    claims := &Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

