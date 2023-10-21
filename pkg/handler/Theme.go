package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func (h *Handler) GetTheme(w http.ResponseWriter, r *http.Request) {
	thema, err := h.services.ThemeService.GetAll()
	if err != nil {
		return
	}

	result, err := json.Marshal(thema)
	if err != nil {
		return
	}

	w.Write(result)
}
func (h *Handler) GetThemeByID(w http.ResponseWriter, r *http.Request) {
	themeID := chi.URLParam(r, "ID")

	thema, err := h.services.ThemeService.GetById(themeID)
	if err != nil {
		return
	}

	result, err := json.Marshal(thema)
	if err != nil {
		return
	}

	w.Write(result)
}
func (h *Handler) CreateTheme(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) UpdateTheme(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DeleteTheme(w http.ResponseWriter, r *http.Request) {

}
