package handler

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/template"
	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
	"github.com/zhashkevych/todo-app/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (h *Handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	// Создание экземпляра структуры
	data, err := h.services.AuthorService.GetAll()

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

func (h *Handler) GetAuthorsByID(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/article.html", "templates/head.html", "templates/header_for_articlePage.html", "templates/footer.html")
	if err != nil {
		log.Print("err :", err.Error())
		return
	}
	authorID := chi.URLParam(r, "ID")
	author, err := h.services.AuthorService.GetById(authorID)
	if err != nil {
		log.Print("err :", err.Error())
		return
	}

	data := struct {
		Author *models.Author
	}{
		Author: author,
	}

	tmpl.ExecuteTemplate(w, "author", data)
}
func (h *Handler) CreateAuthors(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	var authorData models.Author
	if err := json.Unmarshal(body, &authorData); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Создание объекта Author на основе данных AuthorData
	author := models.Author{
		ID:        "test", // генерируем уникальный идентификатор для автора
		FirstName: authorData.FirstName,
		LastName:  authorData.LastName,
		Email:     authorData.Email,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	h.services.AuthorService.Create(author)
	// Отправить ответ клиенту
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Автор успешно создан и сохранен!"))
}

func (h *Handler) UpdateAuthors(w http.ResponseWriter, r *http.Request) {
	// Чтение тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	// Декодирование JSON-данных в объект ArticleData
	var authorData models.Author
	if err := json.Unmarshal(body, &authorData); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Получение ID статьи из URL-параметров
	vars := mux.Vars(r)
	authorID := vars["id"]

	// Получение статьи из базы данных по ее ID
	author, err := h.services.AuthorService.GetById(authorID)
	if err != nil {
		http.Error(w, "Автор не найден", http.StatusNotFound)
		return
	}

	// Обновление полей статьи на основе данных AuthorData
	author = &models.Author{
		ID:        "test", // генерируем уникальный идентификатор для автора
		FirstName: authorData.FirstName,
		LastName:  authorData.LastName,
		Email:     authorData.Email,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	// Сохранение обновленной статьи в базе данных
	author, err = h.services.AuthorService.Update(authorID, *author)
	if err != nil {
		http.Error(w, "Ошибка при обновлении данных автора", http.StatusInternalServerError)
		return
	}

	// Отправка ответа клиенту
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные автора успешно обновлены!"))
}
func (h *Handler) DeleteAuthors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["authorID"]

	// код для удаления статьи с идентификатором articleID из базы данных или другого хранилища данных
	fmt.Fprintf(w, "Автор с идентификатором %s успешно удален", authorID)
}
