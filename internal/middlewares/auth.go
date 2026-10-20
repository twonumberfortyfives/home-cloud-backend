package middlewares

import (
	"log"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, "executing authorization middleware #1")
		next.ServeHTTP(w, r)
		log.Println(r.URL.Path, "executing authorization middleware #2")
	})
}
