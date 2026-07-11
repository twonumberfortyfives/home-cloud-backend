package app

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/twonumberfortyfives/home-cloud-backend/internal/health"
)

func SwaggerInit() http.HandlerFunc {
	return httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/api/swagger/doc.json"))
}

// CheckStatus godoc
//
//	@Summary		Check application health
//	@Description	Returns the health status of the backend
//	@Tags			Health
//	@Produce		plain
//	@Success		200	{string}	string	"OK"
//	@Router			/api/health [get]
func HealthInit() *health.HealthHandler {
	healthService := &health.HealthService{}

	return health.NewHealthHandler(healthService)
}
