package handler

import (
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	_ "github.com/zhashkevych/todo-app/docs"
	"github.com/zhashkevych/todo-app/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", h.MainPage)
	router.Get("/input", h.inputPage)
	router.Post("/save", h.savePage)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/img/*", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	// Создание группы маршрутов для статей
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", h.GetArticles)
		r.Get("/{ID}/{articleTitle}", h.GetArticleByID)
		r.Post("/", h.CreateArticle)
		r.Put("/{articleID}", h.UpdateArticle)
		r.Delete("/{articleID}", h.DeleteArticle)
	})
	return router
}
