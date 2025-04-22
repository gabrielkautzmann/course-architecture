package comments

import (
    "encoding/json"
    "net/http"
    "scalable-go-api/pkg/auth"

    "github.com/gorilla/mux"
)

type CommentPayload struct {
    Content string `json:"content"`
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.ExtractUserID(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var payload CommentPayload
    _ = json.NewDecoder(r.Body).Decode(&payload)

    postID := mux.Vars(r)["post-id"]
    comment, err := CreateComment(postID, userID, payload.Content)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(comment)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
    postID := mux.Vars(r)["post-id"]
    comments, err := GetAllByPostID(postID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(comments)
}
