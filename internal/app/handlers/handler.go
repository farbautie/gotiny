package handlers

import "github.com/farbautie/gotiny/pkg/database/repositories"

type Handler struct {
	Repository *repositories.Repositories
}

func New(rp *repositories.Repositories) *Handler {
	return &Handler{
		Repository: rp,
	}
}
