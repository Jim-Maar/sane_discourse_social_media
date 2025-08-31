package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostComponentSize = int

const PostComponentSizeSmall PostComponentSize = 3
const PostComponentSizeMedium PostComponentSize = 2
const PostComponentSizeLarge PostComponentSize = 1

type HeaderComponentSize = int

const HeaderComponentSizeVerySmall HeaderComponentSize = 4
const HeaderComponentSizeSmall HeaderComponentSize = 3
const HeaderComponentSizeMedium HeaderComponentSize = 2
const HeaderComponentSizeLarge HeaderComponentSize = 1

type DeviderType string

const RegularDevider = "regular"

type Component interface {
	isComponent()
}

type PostComponent struct {
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
	Size   PostComponentSize  `json:"size" bson:"size"`
}

func (PostComponent) isComponent() {}

type HeaderComponent struct {
	Content string              `json:"content" bson:"content"`
	Size    HeaderComponentSize `json:"size" bson:"size"`
}

func (HeaderComponent) isComponent() {}

type PragraphComponent struct {
	Content string `json:"content" bson:"content"`
}

func (PragraphComponent) isComponent() {}

type DeviderComponent struct {
	Type DeviderType `json:"type" bson:"type"`
}

func (DeviderComponent) isComponent() {}

type Userpage struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Components []Component        `json:"components" bson:"components"`
}

func NewUserpage(components []Component, userID primitive.ObjectID) *Userpage {
	return &Userpage{
		ID:         primitive.NewObjectID(),
		UserID:     userID,
		Components: components,
	}
}
