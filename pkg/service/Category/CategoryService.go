package Category

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
)

type CategoryService struct {
	rep    *repository.Repository
	gen    *tools.UUIDStringGenerator
	Logger logger.Interface
}

func NewCategoryService(rep *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) CategoryService {
	return CategoryService{
		rep:    rep,
		gen:    gen,
		Logger: Logger,
	}
}

func (c CategoryService) Create() (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) Update(id string, category models.Category) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) GetById(id string) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) GetAll() ([]*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
