package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"gorm.io/gorm/logger"
)

type AuthorService struct {
	rep    repository.Repository
	gen    string
	Logger logger.Interface
}

func (a AuthorService) Create(author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) Update(id string, author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
