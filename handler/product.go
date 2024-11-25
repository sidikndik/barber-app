package handler

import (
	"barber-app/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewProductHandler(service service.AllService, log *zap.Logger) ProductHadler {
	return ProductHadler{
		Service: service,
		Log:     log,
	}
}

func (productHadler *ProductHadler) Create(ctx *gin.Context) {

}

func (productHadler *ProductHadler) GetAll(ctx *gin.Context) {

}

func (productHadler *ProductHadler) Login(ctx *gin.Context) {

}
