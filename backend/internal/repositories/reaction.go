package repositories

import (
	"context"
	"sane-discourse-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReactionRepository struct {
	client *mongo.Client
}

func NewReactionRepository(client *mongo.Client) *ReactionRepository {
	return &ReactionRepository{
		client: client,
	}
}

func (r *ReactionRepository) collection() *mongo.Collection {
	return r.client.Database("sane_discourse").Collection("reactions")
}

func (r *ReactionRepository) Create(reaction models.Reaction) (*models.Reaction, error) {
	reaction.ID = primitive.NewObjectID()
	result, err := r.collection().InsertOne(context.TODO(), reaction)
	if err != nil {
		return nil, err
	}
	reaction.ID = result.InsertedID.(primitive.ObjectID)
	return &reaction, nil
}

func (r *ReactionRepository) FindByID(id primitive.ObjectID) (*models.Reaction, error) {
	var reaction models.Reaction
	err := r.collection().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&reaction)
	if err != nil {
		return nil, err
	}
	return &reaction, nil
}

func (r *ReactionRepository) FindByUserID(userID primitive.ObjectID) ([]models.Reaction, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var reactions []models.Reaction
	if err = cursor.All(context.TODO(), &reactions); err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionRepository) FindByPostID(postID primitive.ObjectID) ([]models.Reaction, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{"post_id": postID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var reactions []models.Reaction
	if err = cursor.All(context.TODO(), &reactions); err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionRepository) FindByUserIDAndPostID(userID, postID primitive.ObjectID) ([]models.Reaction, error) {
	filter := bson.M{
		"user_id": userID,
		"post_id": postID,
	}
	cursor, err := r.collection().Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var reactions []models.Reaction
	if err = cursor.All(context.TODO(), &reactions); err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionRepository) FindAll() ([]models.Reaction, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var reactions []models.Reaction
	if err = cursor.All(context.TODO(), &reactions); err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionRepository) Delete(userID, postID primitive.ObjectID) error {
	_, err := r.collection().DeleteOne(context.TODO(), bson.M{
		"user_id": userID,
		"post_id": postID,
	})
	return err
}
