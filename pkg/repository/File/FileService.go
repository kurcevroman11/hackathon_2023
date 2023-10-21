package File

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository/Error"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type FileRepositoryImpl struct {
	db     *gorm.DB
	Logger logger.Interface
}

func NewFileRepository(db *gorm.DB, Logger logger.Interface) FileRepositoryImpl {
	return FileRepositoryImpl{
		db:     db,
		Logger: Logger,
	}
}
func (f FileRepositoryImpl) Create(file *models.File) (*models.File, error) {
	tx := f.db.Create(file)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, Error.RecordNotCreate
	}
	return file, nil
}

func (f FileRepositoryImpl) Update(id string, updated models.File) (*models.File, error) {
	// Найти статью по идентификатору
	var existingArticle models.File
	if err := f.db.Where("id = ?", id).First(&existingArticle).Error; err != nil {
		return nil, err
	}

	existingArticle.Name = updated.Name
	existingArticle.Path = updated.Path

	if err := f.db.Save(&existingArticle).Error; err != nil {
		return nil, err
	}

	return &existingArticle, nil
}

func (f FileRepositoryImpl) GetById(id string) (*models.File, error) {
	var dest *models.File
	result := f.db.Where("ID = ?", id).Find(&dest)
	if result.Error != nil {
		return nil, result.Error
	}
	return dest, nil
}

func (f FileRepositoryImpl) GetAll() ([]*models.File, error) {

	var dest []*models.File
	result := f.db.Find(&dest)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, Error.RecordNotCreate
	}
	return dest, nil
}

func (f FileRepositoryImpl) Delete(id string) (bool, error) {
	result := f.db.Where("id = ?", id).Delete(&models.File{})
	if result.Error != nil {
	}
	return false, result.Error
}
