package posts

import (
    "github.com/google/uuid"
)

func CreatePost(title, userID string, topics []string, body string) (*Post, error) {
    post := &Post{
        ID:     uuid.New().String(),
        Title:  title,
        UserID: userID,
        Topics: topics,
        Body:   body,
    }
    return post, Create(post)
}
