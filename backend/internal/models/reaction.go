package models

import (
	"sane-discourse-backend/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reaction struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ReactionType types.ReactionType `json:"reaction_type" bson:"reaction_type"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID       primitive.ObjectID `json:"post_id" bson:"post_id"`
}

func NewReaction(reactionType types.ReactionType, userID, postID primitive.ObjectID) *Reaction {
	return &Reaction{
		ID:           primitive.NewObjectID(),
		ReactionType: reactionType,
		UserID:       userID,
		PostID:       postID,
	}
}
