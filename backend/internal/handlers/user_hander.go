package handlers

import (
	"sane-discourse-backend/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type LoginRequest struct {
	Username string `json:"username" bson:"username"`
}

/*func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.WithField("error", err.Error()).Error("LoginUser: Invalid request body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.WithField("input", loginRequest).Info("LoginUser: Request received")

	user, err := h.userService.LoginUser(loginRequest.Username)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"input": loginRequest,
			"error": err.Error(),
		}).Error("LoginUser: Request failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output := user.ToPublic()
	log.WithFields(map[string]interface{}{
		"input":  loginRequest,
		"output": output,
	}).Info("LoginUser: Request successful")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}*/
