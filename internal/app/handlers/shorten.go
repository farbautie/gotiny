package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type ShortenRequest struct {
	Url string `json:"url"`
}

func (h *Handler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) handleError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"message": message})
}

func (h *Handler) extractShortUrl(r *http.Request) (string, bool) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) >= 4 && parts[1] == "api" && parts[2] == "v1" && parts[3] == "shorten" {
		return parts[4], true
	}
	return "", false
}

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var data ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.handleError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	shortUrl := uuid.New().String()[:8]
	link, err := h.Repository.Save(data.Url, shortUrl)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondJSON(w, http.StatusCreated, link)
}

func (h *Handler) GetShortenUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl, ok := h.extractShortUrl(r)
	if !ok {
		h.handleError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	link, err := h.Repository.GetByShortUrl(shortUrl)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if link == nil {
		h.handleError(w, http.StatusNotFound, "Link not found")
		return
	}

	updatedLink, err := h.Repository.UpdateStats(link.ID)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondJSON(w, http.StatusOK, updatedLink)
}

func (h *Handler) UpdateShortenUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl, ok := h.extractShortUrl(r)
	if !ok {
		h.handleError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var data ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.handleError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	link, err := h.Repository.GetByShortUrl(shortUrl)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if link == nil {
		h.handleError(w, http.StatusNotFound, "Link not found")
		return
	}

	link.ShortUrl = uuid.New().String()[:8]
	link.Url = data.Url
	if err := h.Repository.Update(link.ID, link); err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "Link updated"})
}

func (h *Handler) DeleteShortenUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl, ok := h.extractShortUrl(r)
	if !ok {
		h.handleError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	link, err := h.Repository.GetByShortUrl(shortUrl)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if link == nil {
		h.handleError(w, http.StatusNotFound, "Link not found")
		return
	}

	if err := h.Repository.Delete(link.ID); err != nil {
		h.handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "Link deleted"})
}
