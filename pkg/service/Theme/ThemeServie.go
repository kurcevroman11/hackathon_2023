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
		Id:    t.gen.GenerateUUID(),
		Name:  "Жёлтая тема",
		Color: "#F0F096",
	}
	theme5 := &models.Theme{
		Id:    "fef43f09-9b56-ad41-8c46-5a01b542ce27",
		Name:  "По умолчанию",
		Color: "#FFFFFF",
	}
	theme1 := &models.Theme{
		Id:    t.gen.GenerateUUID(),
		Name:  "Голубая тема",
		Color: "#DCDCDC",
	}
	theme2 := &models.Theme{
		Id:    t.gen.GenerateUUID(),
		Name:  "Красная тема",
		Color: "#DC9696",
	}
	theme3 := &models.Theme{
		Id:    t.gen.GenerateUUID(),
		Name:  "Темная тема",
		Color: "#6B6A66",
	}

	t.rep.ThemeRepository.Create(theme1)
	t.rep.ThemeRepository.Create(theme2)
	t.rep.ThemeRepository.Create(theme3)
	t.rep.ThemeRepository.Create(theme5)
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
