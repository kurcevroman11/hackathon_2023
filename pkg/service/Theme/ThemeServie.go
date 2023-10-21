package Theme

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"
)

type ThemeServiceImpl struct {
	rep    *repository.Repository
	gen    *tools.UUIDStringGenerator
	Logger logger.Interface
}

func NewThemeServiceImpl(rep *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) ThemeServiceImpl {
	return ThemeServiceImpl{
		rep:    rep,
		gen:    gen,
		Logger: Logger,
	}
}

func (t ThemeServiceImpl) Create(theme *models.Theme) (*models.Theme, error) {
	return t.rep.ThemeRepository.Create(theme)
}
func (t ThemeServiceImpl) FakeDate() (*models.Theme, error) {

	theme := &models.Theme{
		Id:   t.gen.GenerateUUID(),
		Name: "Темная тема",
		R:    10,
		G:    0,
		B:    20,
	}
	return t.rep.ThemeRepository.Create(theme)

}

func (t ThemeServiceImpl) Update(id string, author models.Theme) (*models.Theme, error) {
	//TODO implement me
	panic("implement me")
}

func (t ThemeServiceImpl) GetById(id string) (*models.Theme, error) {
	return t.rep.ThemeRepository.GetById(id)
}

func (t ThemeServiceImpl) GetAll() ([]*models.Theme, error) {
	return t.rep.ThemeRepository.GetAll()
}

func (t ThemeServiceImpl) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
