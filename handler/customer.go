package handler

import (
	"barber-app/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewCustomerHandler(service service.AllService, log *zap.Logger) CustomerHadler {
	return CustomerHadler{
		Service: service,
		Log:     log,
	}
}

func (CustomerHadler *CustomerHadler) Create(ctx *gin.Context) {

}

func (CustomerHadler *CustomerHadler) GetAll(ctx *gin.Context) {

}

func (CustomerHadler *CustomerHadler) Login(ctx *gin.Context) {

}
