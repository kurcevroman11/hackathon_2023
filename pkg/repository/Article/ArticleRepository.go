package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ArticleRepository struct {
	db     *gorm.DB
	Logger logger.Interface
}

func NewArticleRepository(db *gorm.DB, gen *string, Logger logger.Interface) ArticleRepository {
	return ArticleRepository{
		db:     db,
		Logger: Logger,
	}
}

func (a ArticleRepository) Create() (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Update(id string, article models.Article) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) GetById(id string) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) GetAll() ([]*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
