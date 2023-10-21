package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/zhashkevych/todo-app/pkg/models"
	"net/http"
)

// CreateAuthor is a handler function that create new authors.
//
// @summary Создание автора
// @description Эндпоинт для создания нового автора.
// @tags Авторы
// @param {object} w - Ответ HTTP
// @param {object} r - Запрос HTTP
// @produces application/json
// @consumes application/json
// @response 201 - Успешное создание автора
// @response 400 - Ошибка при парсинге JSON
// @response 500 - Ошибка при создании автора
func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var input models.Author

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

// GetAllAuthor is a handler function that returns all authors.
//
// @Summary Get all authors
// @Description Returns a list of all authors.
// @Tags Authors
// @Produce json
// @Success 200 {array} Author
// @Failure 500 {string} string "Failed to get all authors"
// @Failure 500 {string} string "Failed to send response"
// @Router /authors [get]
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

// UpdateAuthor is a handler function that updates an author.
//
// @Summary Update an author
// @Description Updates the details of a specific author.
// @Tags Authors
// @Accept json
// @Param authors path string true "Author ID"
// @Param input body Author true "Author object to be updated"
// @Produce json
// @Success 200 {object} Author
// @Failure 400 {string} string "Invalid JSON"
// @Failure 500 {string} string "Failed to update author"
// @Failure 500 {string} string "Failed to send response"
// @Router /authors/{authors} [put]
func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	// Получение значения ID из URL
	authorID := chi.URLParam(r, "authors")

	var input models.Author

	// Распарсите JSON-данные из тела запроса
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		// Обработайте ошибку парсинга JSON
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	out, err := h.services.AuthorService.Update(authorID, input)
	if err != nil {
		// Обработайте ошибку обновления автора
		http.Error(w, "Failed to update author", http.StatusInternalServerError)
		return
	}

	// Отправьте результат клиенту
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(out); err != nil {
		// Обработайте ошибку сериализации JSON
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

// GetByIdAuthor is a handler function that retrieves an author by ID.
//
// @Summary Get author by ID
// @Description Retrieves the details of a specific author by ID.
// @Tags Authors
// @Produce json
// @Param authorID path string true "Author ID"
// @Success 200 {object} Author
// @Failure 404 {string} string "Author Not Found"
// @Failure 500 {string} string "Failed to get author"
// @Failure 500 {string} string "Failed to marshal author data"
// @Router /authors/{authorID} [get]
func (h *Handler) GetByIdAuthor(w http.ResponseWriter, r *http.Request) {
	// Получение ID автора из URL
	authorID := chi.URLParam(r, "authorID")

	// Вызов метода GetById вашего репозитория для получения автора по ID
	author, err := h.services.AuthorService.GetById(authorID)
	if err != nil {
		// Обработка случаев, когда автор не найден или возникает другая ошибка
		if err.Error() == "Not Found" {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get author", http.StatusInternalServerError)
		return
	}

	// Преобразование автора в JSON
	data, err := json.Marshal(author)
	if err != nil {
		http.Error(w, "Failed to marshal author data", http.StatusInternalServerError)
		return
	}

	// Установка application/json заголовка ответа
	w.Header().Set("Content-Type", "application/json")
	// Отправка JSON-данных ответа
	w.Write(data)
}

// DeleteAuthor is a handler function that deletes an author by ID.
//
// @Summary Delete author by ID
// @Description Deletes a specific author by ID.
// @Tags Authors
// @Param authorID path string true "Author ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "Author Not Found"
// @Failure 500 {string} string "Failed to delete author"
// @Router /authors/{authorID} [delete]
func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	// Получение ID автора из URL
	authorID := chi.URLParam(r, "authorID")

	// Вызов метода Delete вашего репозитория для удаления автора по ID
	success, err := h.services.AuthorService.Delete(authorID)
	if err != nil {
		// Обработка случаев, когда автор не найден или возникает другая ошибка
		if err.Error() == "Not Found" {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}

	// Проверка успешности удаления
	if !success {
		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}

	// Отправка статуса No Content в случае успешного удаления
	w.WriteHeader(http.StatusNoContent)
}
