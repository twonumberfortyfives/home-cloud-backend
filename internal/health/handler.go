package health

import "net/http"

type HealthHandler struct {
	service *HealthService
}

func NewHealthHandler(service *HealthService) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) CheckStatus(w http.ResponseWriter, r *http.Request) {
	status := h.service.CheckStatus()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(status))
}
