package service

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/service/Article"
	"github.com/zhashkevych/todo-app/pkg/service/Author"
	"github.com/zhashkevych/todo-app/pkg/service/File"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
	"net/http"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ArticleService interface {
	Create(article *models.Article) (*models.Article, error)
	Update(id string, article models.Article) (*models.Article, error)
	GetById(id string) (*models.Article, error)
	GetAll(filter *models.FilterArticle) ([]*models.Article, error)
	GenerateQRCode(str string, dest *models.Article) error
	Delete(id string) (bool, error)
	GetImage(r *http.Request) (*models.File, error)
	FakeData() (*models.Article, error)
}
type AuthorService interface {
	Create(author models.Author) (*models.Author, error)
	Update(id string, author models.Author) (*models.Author, error)
	GetById(id string) (*models.Author, error)
	GetAll() ([]*models.Author, error)
	Delete(id string) (bool, error)
}

type FileService interface {
	Create(author *models.File) (*models.File, error)
	Update(id string, author models.File) (*models.File, error)
	GetById(id string) (*models.File, error)
	GetAll() ([]*models.File, error)
	Delete(id string) (bool, error)
}

type Service struct {
	AuthorService
	ArticleService
	FileService
}

func NewService(repos *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) *Service {
	return &Service{
		ArticleService: Article.NewArticleService(repos, gen, Logger),
		AuthorService:  Author.NewAuthorService(repos, gen, Logger),
		FileService:    File.NewFileService(repos, gen, Logger),
	}
}
