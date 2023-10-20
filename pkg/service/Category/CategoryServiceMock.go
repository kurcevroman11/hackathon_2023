package Category

import "github.com/zhashkevych/todo-app/pkg/models"

type CategoryServiceMock struct {
}

func NewCategoryServiceMock() CategoryServiceMock {
	return CategoryServiceMock{}
}

func (c CategoryServiceMock) Create() (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryServiceMock) Update(id string, category models.Category) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryServiceMock) GetById(id string) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryServiceMock) GetAll() ([]*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryServiceMock) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
