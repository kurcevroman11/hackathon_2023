package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type AuthorRepositoryMock struct {
}

func (a AuthorRepositoryMock) Create(author models.Author) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) Update(id string, author models.Author) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) GetById(id string) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) GetAll() ([]*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
