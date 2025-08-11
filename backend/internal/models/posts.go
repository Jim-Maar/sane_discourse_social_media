// internal/models/post.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title" validate:"required,max=200"`
	Description  string             `json:"description" bson:"description" validate:"required,max=10000"`
	ThumbnailURL string             `json:"thumbnail_url" bson:"thumbnail_url"`
	SiteName     string             `json:"site_name" bson:"site_name"`
	URL          string             `json:"url" bson:"url"`
	Type         string             `json:"type" bson:"type"`
	Author       string             `json:"author" bson:"author" validate:"required,max=50"`
}

func NewPost(title, description, thumbnailURL, siteName, url, postType, author string) *Post {
	return &Post{
		ID:           primitive.NewObjectID(),
		Title:        title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
		SiteName:     siteName,
		URL:          url,
		Type:         postType,
		Author:       author,
	}
}
