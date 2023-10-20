package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
)

type ArticleRepositoryMock struct {
}

func (a ArticleRepositoryMock) Create() {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) Update(id string, article models.Article) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepositoryMock) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
