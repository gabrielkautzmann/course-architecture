package images

import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
    postID := mux.Vars(r)["post-id"]

    // Mock URL
    image := map[string]string{
        "id":  uuid.New().String(),
        "url": "https://example.com/image.jpg",
        "postId": postID,
    }

    json.NewEncoder(w).Encode(image)
}
