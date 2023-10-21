package Article

import (
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/zhashkevych/todo-app/pkg/models"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"gorm.io/gorm/logger"

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
func (a ArticleService) FakeData() (*models.Article, error) {

	dest := models.Article{
		ID:              a.gen.GenerateUUID(),
		Title:           "<h1>Test_Title</h1>",
		Content:         "Test_Title",
		PublicationDate: "Test_Data",
		AuthorID:        "",
		Author: models.Author{
			ID:        "test",
			FirstName: "Tests",
			LastName:  "Tests",
			Email:     "Tests@mail.com",
			CreateAt:  time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		},
		Categories: nil,
		CreateAt:   time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	}

	return a.rep.ArticleRepository.Create(&dest)
}

func (a ArticleService) GenerateQRCode(desr *models.Article) error {
	// Создаем QR-код из текста
	valueForQr := fmt.Sprint("/")
	png, err := qrcode.Encode(valueForQr, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	encodedFile := base64.StdEncoding.EncodeToString(png)
	dataURI := "data:image/png;base64," + encodedFile

	// Сохраняем Data URI в модели статьи
	desr.QRCode = dataURI
	return nil
}
