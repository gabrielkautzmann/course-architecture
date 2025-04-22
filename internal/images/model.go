package images

type Image struct {
    ID     string `gorm:"primaryKey"`
    PostID string
    URL    string
}
