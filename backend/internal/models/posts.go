// internal/models/post.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title" validate:"required,max=200"`
	Author       string             `json:"author" bson:"author" validate:"required,max=50"`
	ThumbnailURL string             `json:"thumbnail_url" bson:"thumbnail_url"`
}

func NewPost(title, author, thumbnailURL string) *Post {
	return &Post{
		ID:           primitive.NewObjectID(),
		Title:        title,
		Author:       author,
		ThumbnailURL: thumbnailURL,
	}
}
