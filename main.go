package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"uploader/internal/delivery"
	"uploader/internal/usecase"
	"uploader/pkg/uploader/gcs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// read from env vars
	viper.AutomaticEnv()
	port := viper.GetInt("PORT")
	debug := viper.GetBool("DEBUG")
	swaggerEnabled := viper.GetBool("SWAGGER_ENABLED")
	gcsKeyPath := viper.GetString("GCS_KEY_PATH")
	bucketName := viper.GetString("GCS_BUCKET_NAME")

	// init gcs
	client, err := gcs.NewGCS(gcsKeyPath, bucketName)
	if err != nil {
		logrus.WithError(err).Fatal("unable to init gcs")
	}

	// usecase
	uc := usecase.NewUsecase(client)

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
