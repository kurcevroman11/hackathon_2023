package repository

import (
	"fmt"
	"github.com/zhashkevych/todo-app/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config, loggering bool) (*gorm.DB, error) {
	var newLogger logger.Interface
	if loggering {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // Установите свой log.Logger
			logger.Config{
				SlowThreshold: time.Second, // Порог для медленных запросов
				LogLevel:      logger.Info, // Уровень логирования
				Colorful:      true,        // Включить цветной вывод
			},
		)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	for {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

		dbSQL, err := db.DB()
		if err != nil {
			log.Printf("Ошибка соединение с БД: %s", err.Error())
			return nil, err
		}

		err = dbSQL.Ping()
		if err == nil {
			log.Println("Соединение с базой данных установлено!")
			db.AutoMigrate(models.Article{}, models.File{})
			return db, nil
		}

		log.Println("Соединение с базой данных не установлено. Повторная проверка через 5 секунд...")
		time.Sleep(3 * time.Second)
	}

}
