package comments

import "github.com/google/uuid"

func CreateComment(postID, userID, content string) (*Comment, error) {
    comment := &Comment{
        ID:      uuid.New().String(),
        PostID:  postID,
        UserID:  userID,
        Content: content,
    }
    return comment, Create(comment)
}
