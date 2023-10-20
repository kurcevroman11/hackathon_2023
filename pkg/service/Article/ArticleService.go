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

func (a ArticleService) Create(article *models.Article) (*models.Article, error) {
	return a.rep.ArticleRepository.Create(article)
}

func (a ArticleService) Update(id string, article models.Article) (*models.Article, error) {
	return a.rep.ArticleRepository.Update(id, article)
}

func (a ArticleService) GetById(id string) (*models.Article, error) {
	return a.rep.ArticleRepository.GetById(id)
}

func (a ArticleService) GetAll() ([]*models.Article, error) {
	return a.rep.ArticleRepository.GetAll()
}

func (a ArticleService) Delete(id string) (bool, error) {
	return a.rep.ArticleRepository.Delete(id)
}
