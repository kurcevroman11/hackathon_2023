package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"gorm.io/gorm/logger"
)

type ArticleService struct {
	rep    *repository.Repository
	gen    *string
	Logger logger.Interface
}

func NewArticleService(rep *repository.Repository, gen *string, Logger logger.Interface) ArticleService {
	return ArticleService{
		rep:    rep,
		gen:    gen,
		Logger: Logger,
	}
}

func (a ArticleService) Create() (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) Update(id string, article models.Article) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) GetById(id string) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) GetAll() ([]*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
