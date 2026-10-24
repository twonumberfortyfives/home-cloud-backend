package middlewares

import (
	"fmt"
	"net/http"
	"strings"

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

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		email, _ := payload.Claims["email"].(string)
		id, _ := payload.Claims["id"].(string)

		fmt.Println("!!!!!!!!!")

		fmt.Println(email)

		fmt.Println(payload.Claims)

		_ = id

		next.ServeHTTP(w, r)
	})
}
