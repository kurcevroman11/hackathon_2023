package Category

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CategoryRepository struct {
	db     *gorm.DB
	gen    string
	Logger logger.Interface
}

func (c CategoryRepository) Create() (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) Update(id string, category models.Category) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) GetById(id string) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) GetAll() ([]*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
