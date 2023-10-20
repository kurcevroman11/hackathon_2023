package service

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ArticleService interface {
	Create()
	Update(id string, article models.Article)
	GetById(id string)
	GetAll()
	Delete(id string)
}
type AuthorService interface {
	Create(author models.Author)
	Update(id string, author models.Author)
	GetById(id string)
	GetAll()
	Delete(id string)
}
type CategoryService interface {
	Create()
	Update(id string, category models.Category)
	GetById(id string)
	GetAll()
	Delete(id string)
}

type Service struct {
	AuthorService
	CategoryService
	ArticleService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
