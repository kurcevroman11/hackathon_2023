package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"gorm.io/gorm/logger"
)

type ArticleService struct {
	rep    repository.Repository
	gen    string
	Logger logger.Interface
}

func (a ArticleService) Create() {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) Update(id string, article models.Article) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
