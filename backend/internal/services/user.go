package services

import (
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepository     *repositories.UserRepository
	userpageRepository *repositories.UserpageRepository
}

func NewUserService(userRepository *repositories.UserRepository, userpageRepository *repositories.UserpageRepository) *UserService {
	return &UserService{
		userRepository:     userRepository,
		userpageRepository: userpageRepository,
	}
}

func (s *UserService) LoginUser(username string, email string) (*models.User, error) {
	user, _ := s.userRepository.FindByEmail(email)
	if user != nil {
		return user, nil
	}
	user, err := models.NewUser(username, email)
	if err != nil {
		return nil, err
	}
	user, err = s.userRepository.Create(*user)
	if err != nil {
		return nil, err
	}
	userpage := models.NewUserpage([]models.Component{}, user.ID)
	_, err = s.userpageRepository.Create(*userpage)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *UserService) GetCurrentUser(userID primitive.ObjectID) (*models.User, error) {
	return s.userRepository.FindByID(userID)
}
