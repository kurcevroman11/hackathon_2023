package handler

import (
	"encoding/json"
	"github.com/alecthomas/template"
	"github.com/zhashkevych/todo-app/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetArticles @Summary Получить все статьи
// @Description Получение списка всех статей
// @Tags Articles
// @Produce json
// @Success 200 {array} models.Article
// @Router /articles [get]
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	var articleData models.ArticleData
	if err := json.Unmarshal(body, &articleData); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Создание объекта Article на основе данных ArticleData
	article := models.Article{
		Title:           articleData.Title,
		Content:         articleData.Content,
		PublicationDate: time.Now().String(),
		AuthorID:        "test",
		Author:          models.Author{},
		Categories:      nil,
		CreateAt:        time.Now(),
		UpdatedAt:       time.Now(),
		Image:           "",
		QRCode:          "",
		DeletedAt:       nil,
	}

	h.services.ArticleService.Create(&article)
	// Отправить ответ клиенту
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Статья успешно создана и сохранена!"))
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

	articles, err := h.services.ArticleService.GetAll()

	for _, article := range articles {
		h.services.ArticleService.GenerateQRCode(article)
	}
	data := struct {
		Articles []*models.Article
	}{
		Articles: articles,
	}

	tmpl.ExecuteTemplate(w, "index", data)
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
	content := r.FormValue("editor")

	dest := models.Article{
		Title: title, Content: content, AuthorID: "test"}

	h.services.ArticleService.Create(&dest)

	err := h.services.ArticleService.GenerateQRCode(&dest)
	if err != nil {
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "image/svg+xml")

	// Вывод в буфере
	w.Write([]byte(dest.Image))

}
