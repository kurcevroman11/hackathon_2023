package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type ArticleServiceMock struct {
}

func NewArticleServiceMock() ArticleServiceMock {
	return ArticleServiceMock{}
}

func (a ArticleServiceMock) Create() (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleServiceMock) Update(id string, article models.Article) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleServiceMock) GetById(id string) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleServiceMock) GetAll() ([]*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleServiceMock) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
