package services

import (
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) LoginUser(username string) (*models.User, error) {
	user, _ := s.userRepository.FindByUsername(username)
	if user != nil {
		return user, nil // user exists
	}
	user, err := models.NewUser(username)
	if err != nil {
		return user, err
	}
	user, err = s.userRepository.Create(*user)
	return user, err
}
