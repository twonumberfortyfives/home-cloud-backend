package users

import "net/http"

type UserHandler struct {
}

func (h UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Handler) {
}
