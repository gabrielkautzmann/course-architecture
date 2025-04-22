package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "scalable-go-api/api"
    "scalable-go-api/pkg/database"
    "scalable-go-api/internal/users"
    "scalable-go-api/internal/posts"
    "scalable-go-api/internal/comments"
    "scalable-go-api/internal/votes"
    "scalable-go-api/internal/images"
)

func main() {
    database.Connect()

    // Auto migrate all models
    database.DB.AutoMigrate(&users.User{}, &posts.Post{}, &comments.Comment{}, &votes.Vote{}, &images.Image{})

    r := mux.NewRouter()
    api.RegisterRoutes(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting at port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
