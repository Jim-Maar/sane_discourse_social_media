package models

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
}

func NewUser(username string, email string) (*User, error) {
	if len(username) < 3 {
		return nil, errors.New("username too short")
	}
	if len(email) == 0 {
		return nil, errors.New("email is required")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("email is invalid")
	}
	return &User{
		Username: username,
		Email:    email,
	}, nil
}

func (u *User) ToPublic() *UserPublic {
	return &UserPublic{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

type UserPublic struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
}
