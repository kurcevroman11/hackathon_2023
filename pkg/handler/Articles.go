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
	data, err := h.services.ArticleService.GetAll(&models.FilterArticle{Public: false})

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
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	article := models.Article{}

	title := r.FormValue("title")
	article.Title = title
	subtitle := r.FormValue("subtitle")
	article.Subtitle = subtitle
	createAt := r.FormValue("createAt")
	article.PublicationDate = createAt
	content := r.FormValue("content")
	article.Content = content
	public := r.FormValue("public")
	if public == "true" {
		article.Public = true
	} else {
		article.Public = false
	}

	file, err := h.services.ArticleService.GetImage(r)
	article.ImgFile = *file
	if err != nil {
		http.Error(w, "Ошибка при копировании файла", http.StatusInternalServerError)
		return
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

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/Main.html", "templates/head.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		log.Print("err :", err.Error())
		return
	}

	articles, err := h.services.ArticleService.GetAll(&models.FilterArticle{Public: true})

	for _, article := range articles {
		if len(article.Content) > 300 {
			article.Content = article.Content[:300] + "..."
		}
		article.Content = removeImagesFromContent(article.Content)
	}

	var urls []string

	for _, article := range articles {

		titleURL := strings.ReplaceAll(strings.ToLower(article.Title), " ", "-")
		titleURL = strings.Replace(titleURL, "/", "", -1)
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

	thems, err := h.services.ThemeService.GetAll()
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
	w.Write([]byte(dest.Title))
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

func removeImagesFromContent(content string) string {
	// Создайте новый документ goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Println("Error creating goquery document:", err)
		return content
	}

	// Найдите и удалите все теги <img>
	doc.Find("img").Each(func(index int, element *goquery.Selection) {
		element.Remove()
	})

	// Получите очищенный текстовый контент без изображений
	cleanedContent, _ := doc.Html()

	return cleanedContent
}
