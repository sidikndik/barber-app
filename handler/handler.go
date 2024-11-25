package handler

import (
	"barber-app/service"

	"go.uber.org/zap"
)

type AllHandler struct {
	CustomerHandler CustomerHadler
	ProductHadler   ProductHadler
}

func NewAllHandler(service service.AllService, log *zap.Logger) AllHandler {
	return AllHandler{
		CustomerHandler: NewCustomerHandler(service, log),
		ProductHadler:   NewProductHandler(service, log),
	}
}
