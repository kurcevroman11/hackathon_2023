package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type ArticleRepositoryMock struct {
}

func (a ArticleRepositoryMock) Create() (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) Update(id string, article models.Article) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) GetById(id string) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) GetAll() ([]*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
