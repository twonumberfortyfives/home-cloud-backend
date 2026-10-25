package app

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/twonumberfortyfives/home-cloud-backend/internal/health"
	"github.com/twonumberfortyfives/home-cloud-backend/internal/users"
)

func SwaggerInit() http.HandlerFunc {
	return httpSwagger.Handler(httpSwagger.URL("https://localhost:9000/api/swagger/doc.json"))
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
	return health.CreateHealthHandler(healthService)
}

// CheckStatus godoc
//
//	@Summary		Get Current User
//	@Description	Returns the health status of the backend
//	@Tags			Health
//	@Produce		plain
//	@Success		200	{string}	string	"OK"
//	@Router			/api/user [get]
func UserInit() *users.UserHandler {
	userService := &users.UserService{}
	return users.CreateUserHandler(userService)
}
