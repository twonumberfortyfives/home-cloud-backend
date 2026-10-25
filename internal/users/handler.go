package users

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service *UserService
}

func CreateUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetCurrentUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	user := h.service.GetCurrentUser(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to get current user", http.StatusForbidden)
		return
	}
}
