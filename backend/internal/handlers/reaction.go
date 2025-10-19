package handlers

import (
	"sane-discourse-backend/internal/services"
)

type ReactionHandler struct {
	reactionService *services.ReactionService
}

func NewReactionHandler(reactionService *services.ReactionService) *ReactionHandler {
	return &ReactionHandler{
		reactionService: reactionService,
	}
}
