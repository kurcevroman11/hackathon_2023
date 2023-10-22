package main

import (
	"context"
	"github.com/zhashkevych/todo-app/pkg/tools"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/handler"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/service"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	}, true)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	gen := tools.NewGen()
	repos := repository.NewRepository(db, db.Logger)
	services := service.NewService(repos, &gen, db.Logger)

	//services.ThemeService.FakeDate()

	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
