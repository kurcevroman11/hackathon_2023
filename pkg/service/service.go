package service

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/service/Article"
	"github.com/zhashkevych/todo-app/pkg/service/Author"
	"github.com/zhashkevych/todo-app/pkg/service/Category"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ArticleService interface {
	Create() (*models.Article, error)
	Update(id string, article models.Article) (*models.Article, error)
	GetById(id string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Delete(id string) (bool, error)
}
type AuthorService interface {
	Create(author models.Author) (*models.Author, error)
	Update(id string, author models.Author) (*models.Author, error)
	GetById(id string) (*models.Author, error)
	GetAll() ([]*models.Author, error)
	Delete(id string) (bool, error)
}
type CategoryService interface {
	Create() (*models.Category, error)
	Update(id string, category models.Category) (*models.Category, error)
	GetById(id string) (*models.Category, error)
	GetAll() ([]*models.Category, error)
	Delete(id string) (bool, error)
}

type Service struct {
	AuthorService
	CategoryService
	ArticleService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ArticleService:  Article.NewArticleService(repos, nil, nil),
		AuthorService:   Author.NewAuthorService(repos, nil, nil),
		CategoryService: Category.NewCategoryService(repos, nil, nil),
	}
}
