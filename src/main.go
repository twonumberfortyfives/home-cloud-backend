package main

import (
	"net/http"

	_ "github.com/twonumberfortyfives/home-cloud-backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", healthHandler)

	mux.Handle("/api/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/api/swagger/doc.json"),
	))

	http.ListenAndServe(":8000", mux)
}
