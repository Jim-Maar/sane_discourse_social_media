package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

func (u *User) ToPublic() *UserPublic {
	return &UserPublic{
		ID:       u.ID,
		Username: u.Username,
	}
}

type UserPublic struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}
