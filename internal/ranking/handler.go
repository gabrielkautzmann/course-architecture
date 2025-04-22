package ranking

import (
    "encoding/json"
    "net/http"
)

func GetTopPostsHandler(w http.ResponseWriter, r *http.Request) {
    posts := GetTopPosts24h()
    json.NewEncoder(w).Encode(posts)
}
