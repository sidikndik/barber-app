package model

type Variant struct {
	Base
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}

func VariantSeed() []Variant {
	return []Variant{
		{Name: "Merah"},
		{Name: "Biru"},
		{Name: "Hijau"},
		{Name: "Kuning"},
		{Name: "Hitam"},
		{Name: "Putih"},
		{Name: "Abu-abu"},
		{Name: "Cokelat"},
		{Name: "Ungu"},
		{Name: "Pink"},
		{Name: "XS"},
		{Name: "S"},
		{Name: "M"},
		{Name: "L"},
		{Name: "XL"},
		{Name: "XXL"},
		{Name: "3XL"},
		{Name: "4XL"},
		{Name: "5XL"},
		{Name: "One Size"},
	}
}
