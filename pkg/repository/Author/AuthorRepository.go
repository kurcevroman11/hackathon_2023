package Author

import (
	"errors"
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AuthorRepository struct {
	db     *gorm.DB
	Logger logger.Interface
}

func NewAuthorRepository(db *gorm.DB, logger logger.Interface) *AuthorRepository {
	return &AuthorRepository{
		db:     db,
		Logger: logger,
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
	err := a.db.Model(&models.Author{}).Where("id = ?", id).Updates(author).Error
	if err != nil {
		// Обработка ошибки при обновлении данных в базе данных
		a.Logger.Error(nil, "Ошибка при обновлении автора")
		return nil, err
	}

	return &author, nil
}

func (a *AuthorRepository) GetById(id string) (*models.Author, error) {
	var author models.Author

	// Используйте gorm для выполнения запроса и получения автора по ID
	err := a.db.Where("id = ?", id).First(&author).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Если автор с заданным ID не найден, верните ошибку "Not Found"
			return nil, errors.New("Not Found")
		}
		// Если произошла другая ошибка при выполнении запроса, верните ее
		return nil, err
	}

	return &author, nil
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
	// Используйте gorm для удаления автора из базы данных
	err := a.db.Delete(&models.Author{}, "id = ?", id).Error

	if err != nil {
		// Если произошла ошибка при удалении, верните ее
		return false, err
	}

	return true, nil
}
