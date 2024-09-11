package app

import "net/http"

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	return router
}
