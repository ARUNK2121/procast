package utils

import (
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckAndCreateAdmin(db *gorm.DB, cfg config.Config) {
	//check if any rows available at the admin table
	//if there is no existing admin create a super admin
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := cfg.ADMIN_PASSWORD
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		admin := domain.Admin{
			ID:        1,
			Name:      "procast",
			Email:     cfg.ADMIN_EMAIL,
			Password:  string(hashedPassword),
			Previlege: "super_admin",
		}
		db.Create(&admin)
	}
}