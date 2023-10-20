package Article

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ArticleRepository struct {
	db     *gorm.DB
	gen    string
	Logger logger.Interface
}

func (a ArticleRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Update(id string, article models.Article) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) GetById(id string) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
