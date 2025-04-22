package votes

type Vote struct {
    ID        string `gorm:"primaryKey"`
    EntityID  string
    EntityType string // "post" or "comment"
    UserID    string
    Value     int // +1 or -1
}
