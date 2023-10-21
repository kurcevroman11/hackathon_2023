package Theme

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository/Error"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ThemeRepositoryImpl struct {
	db *gorm.DB

	Logger logger.Interface
}

func NewThemeRepositoryImpl(db *gorm.DB, Logger logger.Interface) ThemeRepositoryImpl {
	return ThemeRepositoryImpl{
		db:     db,
		Logger: Logger,
	}
}

func (t ThemeRepositoryImpl) Create(theme *models.Theme) (*models.Theme, error) {
	tx := t.db.Create(theme)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, Error.RecordNotCreate
	}

	return theme, nil
}

func (t ThemeRepositoryImpl) Update(id string, author models.Theme) (*models.Theme, error) {
	//TODO implement me
	panic("implement me")
}

func (t ThemeRepositoryImpl) GetById(id string) (*models.Theme, error) {
	var dest *models.Theme
	result := t.db.Where("ID = ?", id).Find(&dest)
	if result.Error != nil {
		return nil, result.Error
	}
	return dest, nil
}

func (t ThemeRepositoryImpl) GetAll() ([]*models.Theme, error) {
	var dest []*models.Theme
	result := t.db.Find(&dest)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, Error.RecordNotCreate
	}
	return dest, nil
}

func (t ThemeRepositoryImpl) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
