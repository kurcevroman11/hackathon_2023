package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository/Error"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ArticleRepository struct {
	db     *gorm.DB
	Logger logger.Interface
}

func NewArticleRepository(db *gorm.DB, Logger logger.Interface) ArticleRepository {
	return ArticleRepository{
		db:     db,
		Logger: Logger,
	}
}

func (a ArticleRepository) Create(article *models.Article) (*models.Article, error) {
	tx := a.db.Create(article)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, Error.RecordNotCreate
	}

	return article, nil
}

func (a ArticleRepository) Update(id string, updated models.Article) (*models.Article, error) {
	// Найти статью по идентификатору
	var existingArticle models.Article
	if err := a.db.Where("id = ?", id).First(&existingArticle).Error; err != nil {
		return nil, err
	}

	existingArticle.Title = updated.Title
	existingArticle.Content = updated.Content
	existingArticle.PublicationDate = updated.PublicationDate

	if err := a.db.Save(&existingArticle).Error; err != nil {
		return nil, err
	}

	return &existingArticle, nil
}

func (a ArticleRepository) GetById(id string) (*models.Article, error) {
	var dest *models.Article
	result := a.db.Where("ID = ?", id).Find(&dest)
	if result.Error != nil {
		return nil, result.Error
	}
	return dest, nil
}

func (a ArticleRepository) GetAll() ([]*models.Article, error) {
	var dest []*models.Article
	result := a.db.Find(&dest)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, Error.RecordNotCreate
	}
	return dest, nil
}

func (a ArticleRepository) Delete(id string) (bool, error) {
	// Удаляем статью из базы данных по идентификатору
	result := a.db.Where("id = ?", id).Delete(&models.Article{})
	if result.Error != nil {
		return false, result.Error
	}

	// Проверяем, сколько записей было удалено
	if result.RowsAffected == 0 {
		return false, Error.RecordNotCreate // Запись не найдена
	}

	return true, nil // Запись успешно удалена
}
