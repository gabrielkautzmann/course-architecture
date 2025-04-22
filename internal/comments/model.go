package comments

type Comment struct {
    ID        string `gorm:"primaryKey"`
    PostID    string
    UserID    string
    Content   string
    Upvotes   int
    Downvotes int
}
