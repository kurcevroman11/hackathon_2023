package repository

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository/Article"
	"github.com/zhashkevych/todo-app/pkg/repository/Author"
	"github.com/zhashkevych/todo-app/pkg/repository/Theme"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ArticleRepository interface {
	Create(*models.Article) (*models.Article, error)
	Update(id string, article models.Article) (*models.Article, error)
	GetById(id string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Delete(id string) (bool, error)
}
type AuthorRepository interface {
	Create(author models.Author) (*models.Author, error)
	Update(id string, author models.Author) (*models.Author, error)
	GetById(id string) (*models.Author, error)
	GetAll() ([]*models.Author, error)
	Delete(id string) (bool, error)
}

type ThemeRepository interface {
	Create(author *models.Theme) (*models.Theme, error)
	Update(id string, author models.Theme) (*models.Theme, error)
	GetById(id string) (*models.Theme, error)
	GetAll() ([]*models.Theme, error)
	Delete(id string) (bool, error)
}

type Repository struct {
	ArticleRepository
	AuthorRepository
	ThemeRepository
}

func NewRepository(db *gorm.DB, Logger logger.Interface) *Repository {
	return &Repository{
		ArticleRepository: Article.NewArticleRepository(db, Logger),
		AuthorRepository:  Author.NewAuthorRepository(db, Logger),
		ThemeRepository:   Theme.NewThemeRepositoryImpl(db, Logger),
	}

}
