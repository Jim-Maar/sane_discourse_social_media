package repositories

import (
	"context"
	"sane-discourse-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserpageRepository struct {
	client *mongo.Client
}

func NewUserpageRepository(client *mongo.Client) *UserpageRepository {
	return &UserpageRepository{
		client: client,
	}
}

func (r *UserpageRepository) collection() *mongo.Collection {
	return r.client.Database("sane_discourse").Collection("userpages")
}

func (r *UserpageRepository) Create(userpage models.Userpage) (*models.Userpage, error) {
	result, err := r.collection().InsertOne(context.TODO(), userpage)
	if err != nil {
		return nil, err
	}
	userpage.ID = result.InsertedID.(primitive.ObjectID)
	return &userpage, nil
}

func (r *UserpageRepository) FindByID(id primitive.ObjectID) (*models.Userpage, error) {
	var userpage models.Userpage
	err := r.collection().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&userpage)
	if err != nil {
		return nil, err
	}
	return &userpage, nil
}

func (r *UserpageRepository) FindByUserID(userID primitive.ObjectID) (*models.Userpage, error) {
	var userpage models.Userpage
	err := r.collection().FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&userpage)
	if err != nil {
		return nil, err
	}
	return &userpage, nil
}

func (r *UserpageRepository) FindAll() ([]models.Userpage, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var userpages []models.Userpage
	for cursor.Next(context.TODO()) {
		var userpage models.Userpage
		if err := cursor.Decode(&userpage); err != nil {
			return nil, err
		}
		userpages = append(userpages, userpage)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userpages, nil
}

func (r *UserpageRepository) Update(userpage models.Userpage) (*models.Userpage, error) {
	_, err := r.collection().ReplaceOne(context.TODO(), bson.M{"_id": userpage.ID}, userpage)
	if err != nil {
		return nil, err
	}
	return &userpage, nil
}

func (r *UserpageRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection().DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
