package repositories

import (
	"context"
	"sane-discourse-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) collection() *mongo.Collection {
	return r.client.Database("sane_discourse").Collection("users")
}

func (r *UserRepository) Create(user models.User) (*models.User, error) {
	user.ID = primitive.NewObjectID()
	result, err := r.collection().InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (r *UserRepository) FindByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.collection().FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection().FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(user models.User) (*models.User, error) {
	_, err := r.collection().ReplaceOne(context.TODO(), bson.M{"_id": user.ID}, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection().DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
