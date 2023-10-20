package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type AuthorServiceMock struct {
}

func (a AuthorServiceMock) Create(author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) Update(id string, author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a AuthorServiceMock) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
