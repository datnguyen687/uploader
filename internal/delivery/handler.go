package delivery

import (
	"context"
	"fmt"
	"net/http"

	_ "ecommerce/docs"
	"ecommerce/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHttpDelivery(uc usecase.IUsecase) *httpDelivery {
	return &httpDelivery{
		uc: uc,
	}
}

type httpDelivery struct {
	engine *gin.Engine
	uc     usecase.IUsecase
}

func (d *httpDelivery) Init(r *gin.Engine, debug, swaggerEnabled bool) error {
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}
	logrus.Info("initializing ...")

	if swaggerEnabled {
		// use ginSwagger middleware to serve the API docs
		r.GET("/internal/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", d.ping)

		products := v1.Group("/products")
		{
			products.GET("", d.handleProducts)
		}

	}

	d.engine = r

	return nil
}

func (d *httpDelivery) Run(port int) error {
	addr := fmt.Sprintf(":%d", port)
	logrus.WithField("addr", addr).Info("listening ...")

	return d.engine.Run(addr)
}

func (d *httpDelivery) Shutdown(ctx context.Context) error {
	logrus.Info("shutting down ...")
	return nil
}

// @Summary      ping
// @Description  ping
// @Tags         healthcheck
// @Produce      json
// @Success      200  {object}  BasicResponse
// @Router       /v1/ping [get]
func (d *httpDelivery) ping(c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{
		Code:    StatusCodeOk,
		Message: StatusCodeOk.Message(),
		Error:   "",
	})
}
