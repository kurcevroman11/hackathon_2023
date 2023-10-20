package repository

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository/Article"
	"github.com/zhashkevych/todo-app/pkg/repository/Author"
	"github.com/zhashkevych/todo-app/pkg/repository/Category"
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
type CategoryRepository interface {
	Create() (*models.Category, error)
	Update(id string, category models.Category) (*models.Category, error)
	GetById(id string) (*models.Category, error)
	GetAll() ([]*models.Category, error)
	Delete(id string) (bool, error)
}

type Repository struct {
	ArticleRepository
	CategoryRepository
	AuthorRepository
}

func NewRepository(db *gorm.DB, gen *string, Logger logger.Interface) *Repository {
	return &Repository{
		ArticleRepository:  Article.NewArticleRepository(db, gen, Logger),
		CategoryRepository: Category.NewCategoryRepository(db, gen, Logger),
		AuthorRepository:   Author.NewAuthorRepository(db, gen, Logger),
	}

}
