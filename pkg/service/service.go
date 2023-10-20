package service

import (
	"github.com/zhashkevych/todo-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type AuthorService interface {
	Create
}
type CategoryService interface {
}
type ArticleService interface {
}

type Service struct {
	AuthorService
	CategoryService
	ArticleService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
