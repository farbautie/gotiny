package handlers

import "github.com/farbautie/gotiny/pkg/database/repositories"

type Handler struct {
	repository *repositories.Repositories
}

func New(rp *repositories.Repositories) *Handler {
	return &Handler{
		repository: rp,
	}
}
