package users

import (
    "github.com/google/uuid"
    "scalable-go-api/pkg/database"
)

func CreateUser(username, password string) (*User, error) {
    user := User{
        ID:       uuid.New().String(),
        Username: username,
        Password: password, // NOTE: In production, hash the password!
    }
    result := database.DB.Create(&user)
    return &user, result.Error
}
