package services

import (
	"sane-discourse-backend/internal/repositories"
)

type ReactionService struct {
	reactionRepository *repositories.ReactionRepository
}

func NewReactionService(reactionRepository *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{
		reactionRepository: reactionRepository,
	}
}
