package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/zhashkevych/todo-app/pkg/models"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// GetTheme is a handler function that retrieves all themes.
//
// @Summary Retrieve all themes
// @Produce json
// @Success 200 {array} models.Theme
// @Failure 404 {object} ErrorResponse
// @Router /theme/ [get]
func (h *Handler) GetTheme(w http.ResponseWriter, r *http.Request) {
	var thema []*models.Theme
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

// GetThemeByID is a handler function that retrieves a theme by ID.
//
// @Summary Retrieve a theme by ID
// @Produce json
// @Param ID path string true "Theme ID"
// @Success 200 {object} models.Theme
// @Failure 404 {object} ErrorResponse
// @Router /theme/{ID} [get]
func (h *Handler) GetThemeByID(w http.ResponseWriter, r *http.Request) {
	themeID := chi.URLParam(r, "ID")

	theme, err := h.services.ThemeService.GetById(themeID)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(theme)
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
