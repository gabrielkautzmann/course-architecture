package users

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
    "os"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(secret)
}
