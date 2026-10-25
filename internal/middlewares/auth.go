package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/twonumberfortyfives/home-cloud-backend/internal/users"
	"google.golang.org/api/idtoken"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenStr == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		audience := "416600022776-milrjd2pkohrt5gtie8mjq8ugkdkgh22.apps.googleusercontent.com"

		payload, err := idtoken.Validate(r.Context(), tokenStr, audience)

		fmt.Println(err)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		email, _ := payload.Claims["email"].(string)

		user := users.UserDto{
			Email: email,
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
