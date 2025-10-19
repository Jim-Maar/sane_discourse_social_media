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

type Component struct {
	// Only one of these will be non-nil
	Header    *HeaderComponent   `json:"header,omitempty" bson:"header,omitempty"`
	Post      *PostComponent     `json:"post,omitempty" bson:"post,omitempty"`
	Paragraph *PragraphComponent `json:"paragraph,omitempty" bson:"paragraph,omitempty"`
	Divider   *DividerComponent  `json:"divider,omitempty" bson:"divider,omitempty"`
}

type HeaderComponent struct {
	Content string              `json:"content" bson:"content"`
	Size    HeaderComponentSize `json:"size" bson:"size"`
}

type PostComponent struct {
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
	Size   PostComponentSize  `json:"size" bson:"size"`
}

type PragraphComponent struct {
	Content string `json:"content" bson:"content"`
}

type DividerComponent struct {
	Style DeviderType `json:"style" bson:"style"`
}

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
