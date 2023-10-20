package Article

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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
		ID:              "test1",
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

func (a ArticleService) GenerateQRCode(id string) (barcode.Barcode, error) {
	article, err := a.GetById(id)
	if err != nil {
		return nil, err
	}

	// Создаем QR-код из текста
	qrCode, err := qr.Encode(article.Title, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	// Устанавливаем размер QR-кода (по умолчанию 256x256)
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return nil, err
	}

	return qrCode, err

}
