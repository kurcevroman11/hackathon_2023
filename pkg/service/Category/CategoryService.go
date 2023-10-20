package Category

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"gorm.io/gorm/logger"
)

type CategoryService struct {
	rep    repository.Repository
	gen    string
	Logger logger.Interface
}

func (c CategoryService) Create() {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) Update(id string, category models.Category) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
