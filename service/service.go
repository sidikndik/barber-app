package service

import (
	"barber-app/repository"

	"go.uber.org/zap"
)

type AllService struct {
	CustomerService CustomerServiceInterface
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		CustomerService: NewCustomerService(repo, log),
	}
}
