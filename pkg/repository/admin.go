package repository

import (
	"fmt"

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

func (a *AdminRepository) B() {
	fmt.Println("hyy")
}
