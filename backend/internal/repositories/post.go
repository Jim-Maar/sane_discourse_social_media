package repositories

import (
	"context"
	"sane-discourse-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	client *mongo.Client
}

func NewPostRepository(client *mongo.Client) *PostRepository {
	return &PostRepository{
		client: client,
	}
}

func (r *PostRepository) collection() *mongo.Collection {
	return r.client.Database("sane_discourse").Collection("posts")
}

func (r *PostRepository) Create(post models.Post) (*models.Post, error) {
	post.ID = primitive.NewObjectID()
	result, err := r.collection().InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}
	post.ID = result.InsertedID.(primitive.ObjectID)
	return &post, nil
}

func (r *PostRepository) FindByID(id primitive.ObjectID) (*models.Post, error) {
	var post models.Post
	err := r.collection().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindByURL(url string) (*models.Post, error) {
	var post models.Post
	err := r.collection().FindOne(context.TODO(), bson.M{"url": url}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindAll() ([]models.Post, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var posts []models.Post
	if err = cursor.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(id primitive.ObjectID, update bson.M) (*models.Post, error) {
	filter := bson.M{"_id": id}
	_, err := r.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *PostRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection().DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *PostRepository) FindPostsReactedByUser(userID primitive.ObjectID) ([]models.Post, error) {
	pipeline := []bson.M{
		{"$lookup": bson.M{
			"from":         "reactions",
			"localField":   "_id",
			"foreignField": "post_id",
			"as":           "reactions",
		}},
		{"$match": bson.M{
			"reactions.user_id": userID,
		}},
		{"$project": bson.M{
			"reactions": 0,
		}},
	}

	cursor, err := r.collection().Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var posts []models.Post
	if err = cursor.All(context.TODO(), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) FindAllSortedByReactionCount() ([]models.Post, error) {
	pipeline := []bson.M{
		{"$lookup": bson.M{
			"from":         "reactions",
			"localField":   "_id",
			"foreignField": "post_id",
			"as":           "reactions",
		}},
		{"$addFields": bson.M{
			"reaction_count": bson.M{"$size": "$reactions"},
		}},
		{"$sort": bson.M{
			"reaction_count": -1,
		}},
		{"$project": bson.M{
			"reactions": 0,
		}},
	}

	cursor, err := r.collection().Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var posts []models.Post
	if err = cursor.All(context.TODO(), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
