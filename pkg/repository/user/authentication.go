package user_repository

import (
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	"gorm.io/gorm"
)

type authenticationRepository struct {
	DB *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) interfaces.AuthenticationRepository {
	return &authenticationRepository{
		DB: db,
	}
}
