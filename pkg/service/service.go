package service

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/service/Article"
	"github.com/zhashkevych/todo-app/pkg/service/Author"
	"github.com/zhashkevych/todo-app/pkg/service/Theme"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ArticleService interface {
	Create(article *models.Article) (*models.Article, error)
	Update(id string, article models.Article) (*models.Article, error)
	GetById(id string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	GenerateQRCode(str string, dest *models.Article) error
	Delete(id string) (bool, error)
	FakeData() (*models.Article, error)
}
type AuthorService interface {
	Create(author models.Author) (*models.Author, error)
	Update(id string, author models.Author) (*models.Author, error)
	GetById(id string) (*models.Author, error)
	GetAll() ([]*models.Author, error)
	Delete(id string) (bool, error)
}
type ThemeService interface {
	Create(author *models.Theme) (*models.Theme, error)
	Update(id string, author models.Theme) (*models.Theme, error)
	GetById(id string) (*models.Theme, error)
	GetAll() ([]*models.Theme, error)
	Delete(id string) (bool, error)
}

type Service struct {
	AuthorService
	ArticleService
	ThemeService
}

func NewService(repos *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) *Service {
	return &Service{
		ArticleService: Article.NewArticleService(repos, gen, Logger),
		AuthorService:  Author.NewAuthorService(repos, gen, Logger),
		ThemeService:   Theme.NewThemeServiceImpl(repos, gen, Logger),
	}
}
