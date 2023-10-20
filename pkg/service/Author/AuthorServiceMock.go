package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type AuthorServiceMock struct {
}

func NewAuthorServiceMock() AuthorServiceMock {
	return AuthorServiceMock{}
}

func (a AuthorServiceMock) Create(author models.Author) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) Update(id string, author models.Author) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) GetById(id string) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) GetAll() ([]*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
