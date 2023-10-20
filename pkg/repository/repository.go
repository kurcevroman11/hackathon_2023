package repository

import (
	"gorm.io/gorm"
)

type AuthorService interface {
}
type CategoryService interface {
}
type ArticleService interface {
}

type Repository struct {
	AuthorService
	CategoryService
	ArticleService
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
