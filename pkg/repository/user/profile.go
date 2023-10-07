package user_repository

import (
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	"gorm.io/gorm"
)

type profileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) interfaces.ProfileRepository {
	return &profileRepository{
		DB: db,
	}
}
