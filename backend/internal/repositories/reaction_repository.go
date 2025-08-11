package repositories

import "go.mongodb.org/mongo-driver/mongo"

type ReactionRepository struct {
	client *mongo.Client
}

func NewReactionRepository(client *mongo.Client) *ReactionRepository {
	return &ReactionRepository{
		client: client,
	}
}
