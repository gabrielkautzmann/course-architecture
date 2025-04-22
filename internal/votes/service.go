package votes

import "github.com/google/uuid"

func CastVote(entityID, entityType, userID string, value int) (*Vote, error) {
    vote := &Vote{
        ID:         uuid.New().String(),
        EntityID:   entityID,
        EntityType: entityType,
        UserID:     userID,
        Value:      value,
    }
    return vote, SaveVote(vote)
}
