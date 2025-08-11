package services

import (
	"sane-discourse-backend/internal/repositories"
)

type PostService struct {
	postRepository *repositories.PostRepository
}

func NewPostService(postRepository *repositories.PostRepository) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}
