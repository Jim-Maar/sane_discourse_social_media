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

func (r *UserRepository) FindAll() ([]models.User, error) {
	cursor, err := r.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(id primitive.ObjectID, update bson.M) (*models.User, error) {
	filter := bson.M{"_id": id}
	_, err := r.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *UserRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection().DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
