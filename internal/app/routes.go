package app

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	// Health endpoint init
	mux.HandleFunc("/api/health", HealthInit().CheckStatus)

	// Swagger endpoint init
	mux.Handle("/api/swagger/", SwaggerInit())
}
