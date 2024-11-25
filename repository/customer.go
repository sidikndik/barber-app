package repository

import (
	"barber-app/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepoInterface interface {
	Create(customer *model.Customer) error
	GetByCondition(customer model.Customer) (*model.Customer, error)
	GetAll() (*[]model.Customer, error)
}

type CustomerRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewCustomerRepository(db *gorm.DB, log *zap.Logger) CustomerRepoInterface {
	return &CustomerRepository{
		DB:     db,
		Logger: log,
	}
}

func (customerRepo *CustomerRepository) Create(customer *model.Customer) error {
	query := "INSERT INTO customers (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := customerRepo.DB.Raw(query, customer.Name, customer.Email, customer.Phone, customer.Password).Scan(&customer.ID).Error
	return err
}

func (customerRepo *CustomerRepository) GetAll() (*[]model.Customer, error) {
	var customers []model.Customer
	query := "SELECT name, email, phone, password FROM customers"
	rows, err := customerRepo.DB.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.Customer
		if err := customerRepo.DB.ScanRows(rows, &customers); err != nil {
			return nil, err
		}
		customers = append(customers, user)
	}
	return &customers, nil
}

func (CustomerRepository *CustomerRepository) GetByCondition(customer model.Customer) (*model.Customer, error) {
	var customerResult model.Customer
	query := CustomerRepository.DB.Model(&model.Customer{})
	if customer.Email != "" {
		query = query.Where("email = ?", customer.Email)
	}

	if customer.Phone != "" {
		query = query.Where("phone = ?", customer.Phone)
	}

	err := query.First(&customerResult).Error
	if err != nil {
		return nil, err
	}

	return &customerResult, nil
}
