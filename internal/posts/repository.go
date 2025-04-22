package posts

import (
    "scalable-go-api/pkg/database"
)

func Create(post *Post) error {
    return database.DB.Create(post).Error
}

func GetAll() ([]Post, error) {
    var posts []Post
    err := database.DB.Find(&posts).Error
    return posts, err
}

func GetByID(id string) (*Post, error) {
    var post Post
    err := database.DB.First(&post, "id = ?", id).Error
    return &post, err
}

func DeleteByID(id string) error {
    return database.DB.Delete(&Post{}, "id = ?", id).Error
}
