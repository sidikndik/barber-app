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

func (customerHadler *CustomerHadler) Create(ctx *gin.Context) {

}

func (customerHadler *CustomerHadler) GetAll(ctx *gin.Context) {

}

func (customerHadler *CustomerHadler) Login(ctx *gin.Context) {

}
