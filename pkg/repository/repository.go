package repository

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	Create(author models.Author)
	Update(id string, author models.Author)
	GetById(id string)
	GetAll()
	Delete(id string)
}
type CategoryRepository interface {
	Create()
	Update(id string, category models.Category)
	GetById(id string)
	GetAll()
	Delete(id string)
}
type ArticleRepository interface {
	Create()
	Update(id string, article models.Article)
	GetById(id string)
	GetAll()
	Delete(id string)
}

type Repository struct {
	ArticleRepository
	CategoryRepository
	AuthorRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
