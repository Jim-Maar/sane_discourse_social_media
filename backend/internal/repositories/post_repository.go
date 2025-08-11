package repositories

import "go.mongodb.org/mongo-driver/mongo"

type PostRepository struct {
	client *mongo.Client
}

func NewPostRepository(client *mongo.Client) *PostRepository {
	return &PostRepository{
		client: client,
	}
}
