package services

import (
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/repositories"
	"sane-discourse-backend/pkg/types"
	"sane-discourse-backend/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	postRepository     *repositories.PostRepository
	userRepository     *repositories.UserRepository
	reactionRepository *repositories.ReactionRepository
}

func NewPostService(
	postRepo *repositories.PostRepository,
	userRepo *repositories.UserRepository,
	reactionRepo *repositories.ReactionRepository) *PostService {
	return &PostService{
		postRepository:     postRepo,
		userRepository:     userRepo,
		reactionRepository: reactionRepo,
	}
}

func (s *PostService) CreatePost(url string) (*models.Post, error) {
	linkMetadata, err := utils.ScrapeMetadata(url)
	if err != nil {
		return nil, err
	}

	post := models.NewPost(
		linkMetadata.Title,
		linkMetadata.Description,
		linkMetadata.ImageURL,
		linkMetadata.SiteName,
		linkMetadata.URL,
		linkMetadata.Type,
		linkMetadata.Author,
	)

	return post, nil
}

func (s *PostService) CreatePosts(urls []string) ([]*models.Post, error) {
	posts := []*models.Post{}
	for _, url := range urls {
		post, err := s.CreatePost(url)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostService) AddPosts(posts []models.Post, userId primitive.ObjectID) ([]models.Post, error) {
	addedPosts := []models.Post{}
	_, err := s.userRepository.FindByID(userId)
	if err != nil {
		return addedPosts, err
	}
	for _, post := range posts {
		newPost, _ := s.postRepository.FindByURL(post.URL)
		if newPost == nil {
			newPost, err := s.postRepository.Create(post)
			if err != nil {
				return addedPosts, err
			}
			addedPosts = append(addedPosts, *newPost)
		}
		reactions, _ := s.reactionRepository.FindByUserIDAndPostID(userId, newPost.ID)
		if len(reactions) != 0 {
			continue
		}
		reaction := models.NewReaction(
			types.ReactionTypeAgree,
			userId,
			newPost.ID,
		)
		s.reactionRepository.Create(*reaction)
	}
	return addedPosts, nil
}
