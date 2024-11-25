package model

type Product struct {
	Base
	Name           string           `json:"name" gorm:"type:varchar(100);not null"`
	CategoryID     uint             `json:"category_id" gorm:"not null"`
	Category       Category         `json:"category,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price          float32          `json:"price" gorm:"not null"`
	Photo          string           `json:"photo" gorm:"type:varchar(255)"`
	Description    string           `json:"description" gorm:"type:text"`
	Discount       float32          `json:"discount" gorm:"default:0"`
	ProductVariant []ProductVariant `json:"product_variants,omitempty" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductVariant struct {
	Base
	ProductID uint    `json:"product_id" gorm:"not null;index"`
	VariantID uint    `json:"variant_id" gorm:"not null;index"`
	Stock     int     `json:"stock" gorm:"default:0"`
	Product   Product `json:"product,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Variant   Variant `json:"variant,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func ProductSeed() []Product {
	return []Product{
		{Name: "Smartphone", CategoryID: 1, Price: 3000000, Photo: "smartphone.jpg", Description: "Smartphone canggih", Discount: 10},
		{Name: "Laptop", CategoryID: 1, Price: 10000000, Photo: "laptop.jpg", Description: "Laptop untuk produktivitas", Discount: 5},
		{Name: "Kemeja Pria", CategoryID: 2, Price: 150000, Photo: "kemeja_pria.jpg", Description: "Kemeja formal pria", Discount: 20},
		{Name: "Gaun Wanita", CategoryID: 2, Price: 250000, Photo: "gaun_wanita.jpg", Description: "Gaun elegan", Discount: 15},
		{Name: "Vitamin C", CategoryID: 3, Price: 50000, Photo: "vitamin_c.jpg", Description: "Vitamin C untuk daya tahan tubuh", Discount: 0},
		{Name: "Snack Cokelat", CategoryID: 4, Price: 20000, Photo: "snack_cokelat.jpg", Description: "Snack manis", Discount: 0},
		{Name: "Minuman Energi", CategoryID: 4, Price: 15000, Photo: "minuman_energi.jpg", Description: "Energi instan", Discount: 5},
		{Name: "Sepeda", CategoryID: 5, Price: 1500000, Photo: "sepeda.jpg", Description: "Sepeda olahraga", Discount: 10},
		{Name: "Sepatu Olahraga", CategoryID: 5, Price: 500000, Photo: "sepatu_olahraga.jpg", Description: "Sepatu nyaman", Discount: 15},
		{Name: "Treadmill", CategoryID: 5, Price: 5000000, Photo: "treadmill.jpg", Description: "Alat olahraga rumah", Discount: 10},
	}
}

func ProductVariantSeed() []ProductVariant {
	return []ProductVariant{
		{ProductID: 1, VariantID: 1, Stock: 10},
		{ProductID: 1, VariantID: 2, Stock: 15},
		{ProductID: 2, VariantID: 1, Stock: 5},
		{ProductID: 3, VariantID: 4, Stock: 20},
		{ProductID: 3, VariantID: 5, Stock: 30},
		{ProductID: 4, VariantID: 4, Stock: 10},
		{ProductID: 5, VariantID: 3, Stock: 50},
		{ProductID: 6, VariantID: 2, Stock: 40},
		{ProductID: 7, VariantID: 2, Stock: 25},
		{ProductID: 8, VariantID: 3, Stock: 8},
	}
}
