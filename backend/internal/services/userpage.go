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
	// Insert component at index without modifying the slice during operation
	newComponents := make([]models.Component, len(userpage.Components)+1)
	copy(newComponents[:index], userpage.Components[:index])
	newComponents[index] = *component
	copy(newComponents[index+1:], userpage.Components[index:])
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
	// Remove component from prevIndex
	newComponents := make([]models.Component, len(userpage.Components)-1)
	copy(newComponents[:prevIndex], userpage.Components[:prevIndex])
	copy(newComponents[prevIndex:], userpage.Components[prevIndex+1:])

	// Adjust newIndex if necessary
	if newIndex > prevIndex {
		newIndex -= 1
	}

	// Insert component at newIndex
	result := make([]models.Component, len(newComponents)+1)
	copy(result[:newIndex], newComponents[:newIndex])
	result[newIndex] = component
	copy(result[newIndex+1:], newComponents[newIndex:])

	userpage.Components = result
	userpage, err = s.userpageRepository.Update(*userpage)
	if err != nil {
		return nil, err
	}
	return userpage, nil
}
