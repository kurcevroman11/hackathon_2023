package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type AuthorRepositoryMock struct {
}

func (a AuthorRepositoryMock) Create(author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) Update(id string, author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepositoryMock) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
