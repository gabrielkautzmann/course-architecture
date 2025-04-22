package comments

import (
    "scalable-go-api/pkg/database"
)

func Create(comment *Comment) error {
    return database.DB.Create(comment).Error
}

func GetAllByPostID(postID string) ([]Comment, error) {
    var comments []Comment
    err := database.DB.Where("post_id = ?", postID).Find(&comments).Error
    return comments, err
}
