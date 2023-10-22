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
		Title:           "Томат",
		Subtitle:        "Тома́т или помидо́р (лат. Solánum lycopérsicum) — однолетнее или многолетнее травянистое растение, вид рода Паслён (Solanum)[1] семейства Паслёновые (Solanaceae).",
		Content:         "Название[править | править код]\n«Томат» восходит к ацтекскому названию растения «томатль»[2]. В русский язык слово попало из французского (tomate). «Помидор», другой популярный вариант названия овоща, происходит от итал. pomo d'oro — «золотое яблоко»[3].\n\nСуществует также ещё одна версия: вскоре после открытия Америки, ярко-красные помидоры в Европе, видимо, из-за их расцветки, символизирующей любовь, были переименованы в любовные яблоки (фр. pomme d'amour, а по-русски — помидор)[4].\n\nЛатинский суффикс lycopersicum переводится дословно как «волчий персик»[источник не указан 153 дня].\n\nБиологические особенности[править | править код]\n\n\nВ разделе не хватает ссылок на источники (см. рекомендации по поиску).\n\nИнформация должна быть проверяема, иначе она может быть удалена. Вы можете отредактировать статью, добавив ссылки на авторитетные источники в виде сносок. (6 мая 2019)\n\nТомат имеет сильно развитую корневую систему стержневого типа. Корни разветвлённые, растут и формируются быстро. Уходят в землю на большую глубину (при безрассадной культуре до 1 м и более), распространяясь в диаметре на 1,5—2,5 м. При наличии влаги и питания дополнительные корни могут образовываться на любой части стебля, поэтому томат можно размножать не только семенами, но также черенками и боковыми побегами (пасынками). Поставленные в воду, они через несколько суток образуют корни.\n\nСтебель у томата прямостоячий или полегающий, ветвящийся, высотой от 30 см до 2 м и более. Листья непарноперистые, рассечённые на крупные доли, иногда картофельного типа. Цветки мелкие, невзрачные, жёлтые различных оттенков, собраны в кисть. Томат — факультативный самоопылитель: в одном цветке имеются мужские и женские органы.\n\nПлоды — сочные многогнёздные ягоды различной формы (от плоско-округлой до цилиндрической; могут быть мелкими (масса до 50 г), средними (51—100 г) и крупными (свыше 100 г, иногда до 800 г и более). Окраска плодов от бледно-розовой до ярко-красной и малиновой, от белой, светло-зелёной, светло-жёлтой до золотисто-жёлтой.\n\n\nЦветки и листья томата\n \nСпелые плоды на ветке\n \nВнешний вид и разрез плода",
		PublicationDate: "22.10.2023",
		ThemeId:         "",
		Theme: models.Theme{
			Id:    "3e4e5e06-95e6-894c-e540-4cf9c5b9bf29",
			Name:  "Зеленая тема",
			Color: "#C8E6C8",
		},
		Public:    true,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
		FileId:    "",
		ImgFile: models.File{
			Id:        "43c7b163-333a-a90c-f733-e3de0a3a4f81",
			Name:      "Bright_red_tomato_and_cross_section02.jpg",
			Path:      "/img/43c7b163-333a-a90c-f733-e3de0a3a4f81.jpg",
			CreateAt:  time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		QRCode:    "",
		DeletedAt: nil,
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
