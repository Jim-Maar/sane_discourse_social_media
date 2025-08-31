package services

import (
	"sane-discourse-backend/internal/models"
	"sane-discourse-backend/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserpageService struct {
	userpageRepository *repositories.UserpageRepository
}

func NewUserpageService(userpageRepository *repositories.UserpageRepository) *UserpageService {
	return &UserpageService{
		userpageRepository: userpageRepository,
	}
}

func (s *UserpageService) AddComponent(userID primitive.ObjectID, index int, component *models.Component) (*models.Userpage, error) {
	userpage, err := s.userpageRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	newComponents := append(append(userpage.Components[:index], *component), userpage.Components[index:]...)
	userpage.Components = newComponents
	userpage, err = s.userpageRepository.Update(*userpage)
	if err != nil {
		return nil, err
	}
	return userpage, nil
}

func (s *UserpageService) MoveComponent(userID primitive.ObjectID, prevIndex int, newIndex int) (*models.Userpage, error) {
	userpage, err := s.userpageRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	component := userpage.Components[prevIndex]
	newComponents := append(userpage.Components[:prevIndex], userpage.Components[prevIndex+1:]...)
	if newIndex > prevIndex {
		newIndex -= 1
	}
	newComponents = append(append(newComponents[:newIndex], component), newComponents[newIndex:]...)
	userpage.Components = newComponents
	userpage, err = s.userpageRepository.Update(*userpage)
	if err != nil {
		return nil, err
	}
	return userpage, nil
}
