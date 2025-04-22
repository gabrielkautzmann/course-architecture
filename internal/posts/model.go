package posts

type Post struct {
    ID        string   `gorm:"primaryKey"`
    Title     string
    UserID    string
    Topics    []string `gorm:"type:text[]"`
    Body      string
    Upvotes   int
    Downvotes int
}
