package users

import (
    "encoding/json"
    "net/http"
    "scalable-go-api/pkg/database"
    "github.com/google/uuid"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    _ = json.NewDecoder(r.Body).Decode(&creds)

    user := User{
        ID:       uuid.New().String(),
        Username: creds.Username,
        Password: creds.Password, // In prod, hash this!
    }

    result := database.DB.Create(&user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    _ = json.NewDecoder(r.Body).Decode(&creds)

    var user User
    result := database.DB.Where("username = ? AND password = ?", creds.Username, creds.Password).First(&user)
    if result.Error != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := GenerateToken(user.ID)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
