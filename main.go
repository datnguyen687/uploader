package main

import (
	"context"
	"ecommerce/internal/delivery"
	"ecommerce/internal/repository"
	"ecommerce/internal/usecase"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// read from env vars
	viper.AutomaticEnv()
	port := viper.GetInt("PORT")
	debug := viper.GetBool("DEBUG")
	swaggerEnabled := viper.GetBool("SWAGGER_ENABLED")

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetInt("DB_PORT")
	dbUsername := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	// open db connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("unable to open db connection")
		os.Exit(1)
	}

	// repository
	repo := repository.NewRepository(db)
	if err := repo.AutoMigrate(); err != nil {
		logrus.WithError(err).Error("unable to migrate")
		os.Exit(1)
	}

	// usecase
	uc := usecase.NewUsecase(repo)

	// init engine
	engine := gin.Default()
	d := delivery.NewHttpDelivery(uc)
	d.Init(engine, debug, swaggerEnabled)

	// start
	go d.Run(port)

	// graceful shutdown
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := d.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("unable to gracefully shut down server")
	}
}
