package File

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
)

type FileServiceImpl struct {
	rep    *repository.Repository
	gen    *tools.UUIDStringGenerator
	Logger logger.Interface
}

func NewFileService(rep *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) FileServiceImpl {
	return FileServiceImpl{
		rep:    rep,
		gen:    gen,
		Logger: Logger,
	}
}

func (f FileServiceImpl) Create(author *models.File) (*models.File, error) {
	author.Id = f.gen.GenerateUUID()
	return f.rep.FileRepository.Create(author)
}

func (f FileServiceImpl) Update(id string, author models.File) (*models.File, error) {
	return f.rep.FileRepository.Update(id, author)
}

func (f FileServiceImpl) GetById(id string) (*models.File, error) {
	return f.rep.FileRepository.GetById(id)
}

func (f FileServiceImpl) GetAll() ([]*models.File, error) {
	return f.rep.FileRepository.GetAll()
}

func (f FileServiceImpl) Delete(id string) (bool, error) {
	return f.rep.FileRepository.Delete(id)
}
