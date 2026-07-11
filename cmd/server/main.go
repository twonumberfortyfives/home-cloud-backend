package main

import (
	"net/http"

	_ "github.com/twonumberfortyfives/home-cloud-backend/docs"
	"github.com/twonumberfortyfives/home-cloud-backend/internal/app"
)

func main() {
	mux := http.NewServeMux()

	app.RegisterRoutes(mux)

	http.ListenAndServe(":8000", mux)
}
