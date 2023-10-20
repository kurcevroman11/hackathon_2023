package handler

import (
	"encoding/json"
	"github.com/zhashkevych/todo-app/pkg/models"
	"net/http"
	"time"
)

func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request) {
	// Создание экземпляра структуры
	data := models.Article{
		ID:              "12",
		Title:           "12",
		Content:         "12",
		PublicationDate: "23",
		AuthorID:        "34",
		Author:          models.Author{},
		Categories:      nil,
		CreateAt:        time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	}

	// Преобразование структуры в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Обработка ошибки
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON-данные в ответе
	w.Write(jsonData)
}
func (h *Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
