package handler

import (
	"encoding/json"
	"github.com/alecthomas/template"
	"github.com/zhashkevych/todo-app/pkg/models"
	"log"
	"net/http"
)

func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request) {
	// Создание экземпляра структуры
	data, err := h.services.ArticleService.GetAll()

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

func (h *Handler) MainPaig(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/Main.html", "templates/head.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		log.Print("err :", err.Error())
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}
func (h *Handler) InputPaig(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html", "templates/head.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		log.Print("err :", err.Error())
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}
func (h *Handler) savePaig(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")

	dest := models.Article{
		Title: title, Content: content, AuthorID: "test"}

	h.services.ArticleService.Create(&dest)
}
