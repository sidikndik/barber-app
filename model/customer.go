package model

type Customer struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone    string `json:"phone" gorm:"type:varchar(15);unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

func SeedCustomer() []Customer {
	return []Customer{
		{Name: "Andi Pratama", Email: "andi.pratama@example.com", Phone: "081234567890", Password: "password123"},
		{Name: "Siti Nurhaliza", Email: "siti.nurhaliza@example.com", Phone: "081298765432", Password: "sitinur123"},
		{Name: "Budi Santoso", Email: "budi.santoso@example.com", Phone: "085678901234", Password: "budi12345"},
		{Name: "Rina Oktaviani", Email: "rina.oktaviani@example.com", Phone: "082112233445", Password: "rinaokt123"},
		{Name: "Eka Wibowo", Email: "eka.wibowo@example.com", Phone: "081223344556", Password: "ekawibowo123"},
		{Name: "Ahmad Fadli", Email: "ahmad.fadli@example.com", Phone: "087712345678", Password: "ahmadfadli123"},
		{Name: "Dewi Lestari", Email: "dewi.lestari@example.com", Phone: "089812345678", Password: "dewiles123"},
		{Name: "Yusuf Hidayat", Email: "yusuf.hidayat@example.com", Phone: "081334455667", Password: "yusufhid123"},
		{Name: "Lina Marlina", Email: "lina.marlina@example.com", Phone: "085723456789", Password: "linamar123"},
		{Name: "Agus Setiawan", Email: "agus.setiawan@example.com", Phone: "083312345678", Password: "agusset123"},
	}
}
