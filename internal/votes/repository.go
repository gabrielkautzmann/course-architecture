package votes

import (
    "scalable-go-api/pkg/database"
)

func SaveVote(vote *Vote) error {
    return database.DB.Create(vote).Error
}
