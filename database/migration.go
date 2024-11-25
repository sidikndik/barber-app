package database

import (
	"barber-app/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Customer{},
		&model.Category{},
		&model.Variant{},
		&model.Product{},
		&model.ProductVariant{},
		&model.Order{},
		&model.OrderDetail{},
	)
}
