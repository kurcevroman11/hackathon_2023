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

func (a AuthorRepository) Create(author models.Author) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) Update(id string, author models.Author) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) GetById(id string) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) GetAll() ([]*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
