package handler

import (
	"encoding/json"
	"github.com/zhashkevych/todo-app/pkg/models"
	"net/http"
	"time"
)

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var input models.Author

	genTime := time.Now().Format("2006-01-02 15:04:05.999999999 -0700")
	Time, _ := time.Parse("2006-01-02 15:04:05.99999 -0700", genTime)

	input.CreateAt = Time

	// Распарсите JSON-данные из тела запроса
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		// Обработайте ошибку парсинга JSON
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	out, err := h.services.AuthorService.Create(input)
	if err != nil {
		// Обработайте ошибку создания автора
		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}

	// Отправьте результат клиенту
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(out); err != nil {
		// Обработайте ошибку сериализации JSON
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}
func (h *Handler) GetAllAuthor(w http.ResponseWriter, r *http.Request) {

	out, err := h.services.AuthorService.GetAll()
	if err != nil {
		// Обработайте ошибку создания автора
		http.Error(w, "Failed to get all author", http.StatusInternalServerError)
		return
	}

	// Отправьте результат клиенту
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(out); err != nil {
		// Обработайте ошибку сериализации JSON
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}
