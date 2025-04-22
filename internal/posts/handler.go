package posts

import (
    "encoding/json"
    "net/http"
    "scalable-go-api/internal/users"
    "scalable-go-api/pkg/auth"
    "strings"

    "github.com/gorilla/mux"
)

type PostPayload struct {
    Title  string   `json:"title"`
    Body   string   `json:"body"`
    Topics []string `json:"topics"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.ExtractUserID(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var payload PostPayload
    _ = json.NewDecoder(r.Body).Decode(&payload)

    post, err := CreatePost(payload.Title, userID, payload.Topics, payload.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(post)
}

func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
    posts, err := GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(posts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    post, err := GetByID(vars["post-id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    err := DeleteByID(vars["post-id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
