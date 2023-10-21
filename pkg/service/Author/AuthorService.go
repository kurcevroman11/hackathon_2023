package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
)

type AuthorService struct {
	rep    *repository.Repository
	gen    *tools.UUIDStringGenerator
	Logger logger.Interface
}

func NewAuthorService(rep *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) AuthorService {
	return AuthorService{
		rep:    rep,
		gen:    gen,
		Logger: Logger,
	}
}

func (a AuthorService) Create(author models.Author) (*models.Author, error) {
	author.ID = a.gen.GenerateUUID()
	return a.rep.AuthorRepository.Create(author)
}

func (a AuthorService) Update(id string, author models.Author) (*models.Author, error) {
	return a.rep.AuthorRepository.Update(id, author)
}

func (a AuthorService) GetById(id string) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) GetAll() ([]*models.Author, error) {
	return a.rep.AuthorRepository.GetAll()
}

func (a AuthorService) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
