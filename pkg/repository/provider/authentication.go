package providerRepository

import "gorm.io/gorm"

type authenticationRepository struct {
	DB *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *authenticationRepository {
	return &authenticationRepository{
		DB: db,
	}
}
