package votes

import (
    "encoding/json"
    "net/http"
    "scalable-go-api/pkg/auth"
    "github.com/gorilla/mux"
)

type VotePayload struct {
    Value int `json:"value"` // 1 or -1
}

func VotePostHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.ExtractUserID(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var payload VotePayload
    _ = json.NewDecoder(r.Body).Decode(&payload)

    postID := mux.Vars(r)["post-id"]
    vote, err := CastVote(postID, "post", userID, payload.Value)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(vote)
}
