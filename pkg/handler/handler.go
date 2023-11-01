package handler

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

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
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Разрешить запросы от всех источников
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	router := chi.NewRouter()
	router.Use(corsMiddleware.Handler)
	router.Get("/", h.MainPage)
	router.Get("/input", h.inputPage)
	router.Post("/save", h.savePage)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost/swagger/doc.json"), //The url pointing to API definition
	))

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/img/*", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	// Создание группы маршрутов для статей
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", h.GetArticles)
		r.Get("/{ID}/{articleTitle}", h.GetArticleByID)
		r.Get("/{ID}", h.GetArticleByIDItem)
		r.Post("/", h.CreateArticle)
		r.Put("/{ID}", h.UpdateArticle)
		r.Delete("/{ID}", h.DeleteArticle)
	})
	router.Route("/theme", func(r chi.Router) {
		r.Get("/", h.GetTheme)
		r.Get("/{ID}", h.GetThemeByID)
	})
	return router
}
