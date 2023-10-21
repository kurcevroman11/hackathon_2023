package handler

import (
	"encoding/json"
	"github.com/alecthomas/template"
	"github.com/go-chi/chi"
	"github.com/zhashkevych/todo-app/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

// GetArticleByID is a handler function that retrieves an article by its ID.
//
// @Summary Получение статьи по ID
// @Description Эндпоинт для получения статьи по указанному ID.
// @Tags Статьи
// @Accept  json
// @Produce  json
// @Param ID path string true "ID статьи"
// @Success 200 {object} models.Article
// @Failure 404 {string} string "Статья не найдена"
// @Router /articles/{ID}/{articleTitle} [get]
func (h *Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/article.html", "templates/head.html", "templates/header_for_articlePage.html", "templates/footer.html")
	if err != nil {
		log.Print("err :", err.Error())
		return
	}
	articleId := chi.URLParam(r, "ID")
	article, err := h.services.ArticleService.GetById(articleId)
	if err != nil {
		log.Print("err :", err.Error())
		return
	}
	h.services.ArticleService.GenerateQRCode(r.URL.String(), article)

	data := struct {
		Article *models.Article
	}{
		Article: article,
	}

	tmpl.ExecuteTemplate(w, "articles", data)
}

// CreateArticle is a handler function that creates a new article.
//
// @Summary Создание статьи
// @Description Эндпоинт для создания новой статьи.
// @Tags Статьи
// @Accept json
// @Produce plain
// @Param body body models.ArticleData true "Данные статьи"
// @Success 201 {string} string "Статья успешно создана и сохранена!"
// @Failure 400 {string} string "Ошибка чтения тела запроса | Ошибка декодирования JSON"
// @Router /articles/ [post]
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

// UpdateArticle is a handler function that updates an existing article.
//
// @Summary Обновление статьи
// @Description Эндпоинт для обновления существующей статьи.
// @Tags Статьи
// @Accept json
// @Produce plain
// @Param ID path string true "ID статьи"
// @Param body body models.ArticleData true "Данные статьи"
// @Success 200 {string} string "Статья успешно обновлена!"
// @Failure 400 {string} string "Ошибка чтения тела запроса | Ошибка декодирования JSON"
// @Failure 404 {string} string "Статья не найдена"
// @Failure 500 {string} string "Ошибка при обновлении статьи"
// @Router /articles/{ID} [put]
func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// Чтение тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	// Декодирование JSON-данных в объект ArticleData
	var articleData models.ArticleData
	if err := json.Unmarshal(body, &articleData); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Получение ID статьи из URL-параметров
	articleID := chi.URLParam(r, "articles")

	// Получение статьи из базы данных по ее ID
	article, err := h.services.ArticleService.GetById(articleID)
	if err != nil {
		http.Error(w, "Статья не найдена", http.StatusNotFound)
		return
	}

	// Обновление полей статьи на основе данных ArticleData
	article.Title = articleData.Title
	article.Content = articleData.Content
	article.UpdatedAt = time.Now()

	// Сохранение обновленной статьи в базе данных
	article, err = h.services.ArticleService.Update(articleID, *article)
	if err != nil {
		http.Error(w, "Ошибка при обновлении статьи", http.StatusInternalServerError)
		return
	}

	// Отправка ответа клиенту
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Статья успешно обновлена!"))
}
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/Main.html", "templates/head.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		log.Print("err :", err.Error())
		return
	}

	articles, err := h.services.ArticleService.GetAll()

	for _, article := range articles {
		if len(article.Content) > 900 {
			article.Content = article.Content[:900] + "..."
		}
	}

	var urls []string

	for _, article := range articles {

		titleURL := strings.ReplaceAll(strings.ToLower(article.Title), " ", "-")
		titleURL = Transliterate(titleURL)

		// Форматирование даты публикации в строку "2006-01-02"
		dateStr := article.CreateAt.Format("2006-01-02")
		// Конкатенация заголовка и даты для создания URL
		articleURL := article.ID + "/" + titleURL + "-" + dateStr
		urls = append(urls, articleURL)

	}

	data := struct {
		Articles []*models.Article
		Urls     []string
	}{
		Articles: articles,
		Urls:     urls,
	}

	tmpl.ExecuteTemplate(w, "index", data)
}
func (h *Handler) inputPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html", "templates/head.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		log.Print("err :", err.Error())
		return
	}

	thems := []models.Theme{{
		Id:   1,
		Name: "Темная тема",
		R:    0,
		G:    0,
		B:    0,
	}, {
		Id:   1,
		Name: "Красная",
		R:    255,
		G:    0,
		B:    0,
	},
	}

	if err != nil {
		log.Print("err :", err.Error())
		return
	}

	tmpl.ExecuteTemplate(w, "index", thems)
}
func (h *Handler) savePage(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("editor")

	dest := models.Article{
		Title: title, Content: content, AuthorID: "test"}

	if title == "" || content == "" {
		http.Error(w, "Title and content cannot be empty", http.StatusBadRequest)
		return
	}
	h.services.ArticleService.Create(&dest)

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "image/svg+xml")

	// Вывод в буфере
	w.Write([]byte(dest.Image))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Transliterate(input string) string {
	// Создаем словарь для замены символов
	translit := map[rune]string{
		'А': "A", 'Б': "B", 'В': "V", 'Г': "G", 'Д': "D", 'Е': "E", 'Ё': "YO",
		'Ж': "ZH", 'З': "Z", 'И': "I", 'Й': "Y", 'К': "K", 'Л': "L", 'М': "M",
		'Н': "N", 'О': "O", 'П': "P", 'Р': "R", 'С': "S", 'Т': "T", 'У': "U",
		'Ф': "F", 'Х': "KH", 'Ц': "TS", 'Ч': "CH", 'Ш': "SH", 'Щ': "SCH",
		'Ъ': "", 'Ы': "Y", 'Ь': "", 'Э': "E", 'Ю': "YU", 'Я': "YA",
		'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "yo",
		'ж': "zh", 'з': "z", 'и': "i", 'й': "y", 'к': "k", 'л': "l", 'м': "m",
		'н': "n", 'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u",
		'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch", 'ш': "sh", 'щ': "sch",
		'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "yu", 'я': "ya",
	}

	var result string

	for _, char := range input {
		if replacement, found := translit[char]; found {
			result += replacement
		} else {
			result += string(char)
		}
	}

	return result
}
