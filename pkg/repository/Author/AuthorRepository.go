package Author

import (
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AuthorRepository struct {
	db     *gorm.DB
	Logger logger.Interface
}

func NewAuthorRepository(db *gorm.DB, Logger logger.Interface) AuthorRepository {
	return AuthorRepository{
		db:     db,
		Logger: Logger,
	}
}

func (a AuthorRepository) Create(author models.Author) (*models.Author, error) {
	err := a.db.Create(&author).Error
	if err != nil {
		// Обработка ошибки при сохранении данных в базе данных
		a.Logger.Error(nil, "Ошибка при создании автора: ")
		return nil, err
	}

	return &author, nil
}

func (a AuthorRepository) Update(id string, author models.Author) (*models.Author, error) {
	err := a.db.Create(&author).Error
	if err != nil {
		// Обработка ошибки при сохранении данных в базе данных
		a.Logger.Error(nil, "Ошибка при создании автора: ")
		return nil, err
	}

	return &author, nil
}

func (a AuthorRepository) GetById(id string) (*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorRepository) GetAll() ([]*models.Author, error) {
	var allAuthors []*models.Author

	err := a.db.Find(&allAuthors).Error
	if err != nil {
		a.Logger.Error(nil, "Ошибка при получении всех авторов: ")
		return nil, err
	}
	return allAuthors, nil
}

func (a AuthorRepository) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
