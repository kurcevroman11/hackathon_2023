package handler

import (
	"github.com/go-chi/chi"
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

	router.Get("/", h.MainPaig)

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Создание группы маршрутов для категорий
	router.Route("/categories", func(r chi.Router) {
		r.Get("/", h.GetCategories)
		r.Get("/{categoryID}", h.GetCategoryByID)
		r.Post("/", h.CreateCategory)
		r.Put("/{categoryID}", h.UpdateCategory)
		r.Delete("/{categoryID}", h.DeleteCategory)
	})

	// Создание группы маршрутов для статей
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", h.GetArticles)
		r.Get("/{articleID}", h.GetArticleByID)
		r.Post("/", h.CreateArticle)
		r.Put("/{articleID}", h.UpdateArticle)
		r.Delete("/{articleID}", h.DeleteArticle)
	})

	// Создание группы маршрутов для статей
	router.Route("/authors", func(r chi.Router) {
		r.Get("/", h.GetArticles)
		r.Get("/{authorID}", h.GetArticleByID)
		r.Post("/", h.CreateAuthor)
		r.Put("/{authorID}", h.UpdateArticle)
		r.Delete("/{authorID}", h.DeleteArticle)
	})
	return router
}
