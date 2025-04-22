package users

import (
    "scalable-go-api/pkg/database"
)

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
    var user User
    result := database.DB.Where("username = ? AND password = ?", username, password).First(&user)
    return &user, result.Error
}
