package Category

import "github.com/zhashkevych/todo-app/pkg/models"

type CategoryRepositoryMock struct {
}

func (c CategoryRepositoryMock) Create() (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepositoryMock) Update(id string, category models.Category) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepositoryMock) GetById(id string) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepositoryMock) GetAll() ([]*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepositoryMock) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
