package utils

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckAndCreateAdmin(db *gorm.DB) {
	//check if any contents available at the admin table
	//if there is no existing admin create a super admin
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "procastadminpassword"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		admin := domain.Admin{
			ID:        1,
			Name:      "jerseyhub",
			Email:     "jerseyhub@gmail.com",
			Password:  string(hashedPassword),
			Previlege: "SUPER_ADMIN",
		}
		db.Create(&admin)
	}
}
