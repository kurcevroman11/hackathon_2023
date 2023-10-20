package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AuthorRepository struct {
	db     *gorm.DB
	gen    string
	Logger logger.Interface
}

func (a AuthorRepository) Create(author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) Update(id string, author models.Author) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
