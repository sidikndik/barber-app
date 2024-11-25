package model

type Category struct {
	Base
	Name string `json:"name" gorm:"type:varchar(100);not null"`
}

func CategorySeed() []Category {
	return []Category{
		{Name: "Elektronik"},
		{Name: "Peralatan Rumah Tangga"},
		{Name: "Fashion Pria"},
		{Name: "Fashion Wanita"},
		{Name: "Kesehatan dan Kecantikan"},
		{Name: "Makanan dan Minuman"},
		{Name: "Olahraga dan Outdoor"},
		{Name: "Mainan dan Hobi"},
		{Name: "Buku dan Alat Tulis"},
		{Name: "Otomotif"},
	}
}
