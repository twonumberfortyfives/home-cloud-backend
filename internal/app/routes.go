package app

import (
	"net/http"

	"github.com/twonumberfortyfives/home-cloud-backend/internal/middlewares"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Health endpoint init
	mux.Handle(
		"GET /api/health",
		middlewares.AuthorizationMiddleware(
			http.HandlerFunc(HealthInit().CheckStatus),
		),
	)

	mux.Handle(
		"GET /api/user",
		middlewares.AuthorizationMiddleware(
			http.HandlerFunc(UserInit().GetCurrentUser),
		),
	)

	// Swagger endpoint init
	mux.Handle("/api/swagger/", SwaggerInit())
}
