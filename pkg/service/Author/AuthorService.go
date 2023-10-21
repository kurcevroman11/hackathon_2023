package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
	"time"
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
	genTime := time.Now().Format("2006-01-02 15:04:05.999999999 -0700")
	Time, _ := time.Parse("2006-01-02 15:04:05.99999 -0700", genTime)

	author.CreateAt = Time
	return a.rep.AuthorRepository.Create(author)
}

func (a AuthorService) Update(id string, author models.Author) (*models.Author, error) {
	genTime := time.Now().Format("2006-01-02 15:04:05.999999999 -0700")
	Time, _ := time.Parse("2006-01-02 15:04:05.99999 -0700", genTime)

	author.UpdatedAt = Time
	return a.rep.AuthorRepository.Update(id, author)
}

func (a AuthorService) GetById(id string) (*models.Author, error) {
	return a.rep.AuthorRepository.GetById(id)
}

func (a AuthorService) GetAll() ([]*models.Author, error) {
	return a.rep.AuthorRepository.GetAll()
}

func (a AuthorService) Delete(id string) (bool, error) {
	return a.rep.AuthorRepository.Delete(id)
}
