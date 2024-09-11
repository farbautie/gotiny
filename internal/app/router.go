package app

import (
	"net/http"

	"github.com/farbautie/gotiny/internal/app/handlers"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	handlers := handlers.New()

	router.HandleFunc("GET /api/v1/shorten/{shorten_url}", handlers.GetShortenUrl)
	router.HandleFunc("POST /api/v1/shorten", handlers.ShortenUrl)
	router.HandleFunc("PUT /api/v1/shorten/{shorten_url}", handlers.UpdateShortenUrl)
	router.HandleFunc("DELETE /api/v1/shorten/{shorten_url}", handlers.DeleteShortenUrl)

	return router
}
