package Article

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"time"
)

type ArticleService struct {
	rep    *repository.Repository
	gen    *tools.UUIDStringGenerator
	Logger logger.Interface
}

func NewArticleService(rep *repository.Repository, gen *tools.UUIDStringGenerator, Logger logger.Interface) ArticleService {
	return ArticleService{
		rep:    rep,
		gen:    gen,
		Logger: Logger,
	}
}

func (a ArticleService) Create(article *models.Article) (*models.Article, error) {
	article.ID = a.gen.GenerateUUID()
	article.CreateAt = time.Now()
	article.UpdatedAt = time.Now()
	return a.rep.ArticleRepository.Create(article)
}

func (a ArticleService) Update(id string, article models.Article) (*models.Article, error) {
	return a.rep.ArticleRepository.Update(id, article)
}

func (a ArticleService) GetById(id string) (*models.Article, error) {
	article, err := a.rep.ArticleRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	thems, err := a.rep.ThemeRepository.GetById(article.ThemeId)
	if err != nil {
		return nil, err
	}
	article.Theme = *thems
	file, err := a.rep.FileRepository.GetById(article.FileId)
	if err != nil {
		return nil, err
	}
	article.ImgFile = *file
	return article, err

}

func (a ArticleService) GetAll(filter *models.FilterArticle) ([]*models.Article, error) {
	conditions := a.getWhereList(filter)
	return a.rep.ArticleRepository.GetAll(conditions)
}

func (a ArticleService) Delete(id string) (bool, error) {
	return a.rep.ArticleRepository.Delete(id)
}
func (a ArticleService) FakeData() (*models.Article, error) {

	dest := models.Article{
		ID:              a.gen.GenerateUUID(),
		Title:           "<h1>Test_Title</h1>",
		Content:         "Test_Title",
		PublicationDate: "Test_Data",
		CreateAt:        time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	}

	return a.rep.ArticleRepository.Create(&dest)
}

func (a ArticleService) GenerateQRCode(str string, dest *models.Article) error {
	// Создаем QR-код из текста
	valueForQr := fmt.Sprint(str)
	png, err := qrcode.Encode(valueForQr, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	encodedFile := base64.StdEncoding.EncodeToString(png)
	dataURI := "data:image/png;base64," + encodedFile

	// Сохраняем Data URI в модели статьи
	dest.QRCode = dataURI
	return nil
}

func (a ArticleService) GetImage(r *http.Request) (*models.File, error) {
	maxFileSize := 10 * 1024 * 1024 // Например, 10 МБ

	r.ParseMultipartForm(int64(maxFileSize))
	file, handler, err := r.FormFile("image")
	if err != nil {
		return nil, errors.New("Не удается получить файл")
	}
	defer file.Close()

	// Получаем расширение файла
	fileExt := filepath.Ext(handler.Filename)
	id := a.gen.GenerateUUID()
	// Создаем новый файл для сохранения изображения на сервере
	path := "./img/"
	newFileName := path + id + fileExt
	newFileName2 := "/img/" + id + fileExt
	newFile, err := os.Create(newFileName)
	if err != nil {
		return nil, errors.New("Не удается создать файл для сохранения изображения")
	}
	defer newFile.Close()

	// Копируем содержимое файла в созданный файл
	_, err = io.Copy(newFile, file)
	if err != nil {
		return nil, errors.New("Ошибка при копировании файла")
	}
	localfile := &models.File{
		Id:        id,
		Name:      handler.Filename,
		Path:      newFileName2,
		DeletedAt: nil,
	}
	// В этом моменте, изображение сохранено на сервере с именем newFileName
	a.Logger.Info(context.Background(), "Изображение успешно загружено и сохранено как :", newFileName)
	return localfile, nil

}

func (a ArticleService) getWhereList(filter *models.FilterArticle) []clause.Expression {
	var conditions []clause.Expression

	// Функция для добавления условия в список
	addCondition := func(dataInterface []interface{}, nameField string) {
		conditions = append(conditions, clause.IN{
			Column: clause.Column{Name: nameField},
			Values: dataInterface,
		})
	}

	if filter != nil {
		addCondition([]interface{}{filter.Public}, "public")

	}
	return conditions
}
