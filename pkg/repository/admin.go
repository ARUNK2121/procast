package repository

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

func (a *AdminRepository) GetAdminDetailsByEmail(email string) (domain.Admin, error) {
	var model domain.Admin
	err := a.DB.Raw("SELECT * FROM admins WHERE email = $1", email).Scan(&model).Error
	if err != nil {
		return domain.Admin{}, err
	}

	return model, nil
}
