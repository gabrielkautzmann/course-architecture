package api

import (
    "github.com/gorilla/mux"
    "scalable-go-api/internal/users"
    "scalable-go-api/internal/posts"
    "scalable-go-api/internal/comments"
    "scalable-go-api/internal/images"
    "scalable-go-api/internal/votes"
    "scalable-go-api/internal/ranking"
)

func RegisterRoutes(r *mux.Router) {
    // Users
    r.HandleFunc("/users/create", users.CreateUserHandler).Methods("POST")
    r.HandleFunc("/users/login", users.LoginUserHandler).Methods("POST")

    // Posts
    r.HandleFunc("/posts", posts.CreatePostHandler).Methods("POST")
    r.HandleFunc("/posts", posts.GetAllPostsHandler).Methods("GET")
    r.HandleFunc("/posts/{post-id}", posts.GetPostHandler).Methods("GET")
    r.HandleFunc("/posts/{post-id}", posts.DeletePostHandler).Methods("DELETE")

    // Comments
    r.HandleFunc("/posts/{post-id}/comments", comments.CreateCommentHandler).Methods("POST")
    r.HandleFunc("/posts/{post-id}/comments", comments.GetCommentsHandler).Methods("GET")

    // Images (mock)
    r.HandleFunc("/posts/{post-id}/images", images.UploadImageHandler).Methods("POST")

    // Votes
    r.HandleFunc("/posts/{post-id}/vote", votes.VotePostHandler).Methods("POST")

    // Rankings
    r.HandleFunc("/ranking/top", ranking.GetTopPostsHandler).Methods("GET")
}
