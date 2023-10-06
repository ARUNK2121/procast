package user_repository

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
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

func (a *authenticationRepository) CheckIfPhoneNumberAlreadyExists(phone string) (bool, error) {
	var count int64
	if err := a.DB.Raw("SELECT COUNT(*) FROM users WHERE phone = ?", phone).Scan(&count).Error; err != nil {
		return true, err
	}

	return count > 0, nil

}

func (a *authenticationRepository) UserSignup(model models.UserSignup) error {
	err := a.DB.Exec("INSERT INTO users(name,email,password,phone)VALUES($1,$2,$3,$4)", model.Name, model.Email, model.Password, model.Phone).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *authenticationRepository) CheckIfUserExistsByUsername(username string) (bool, error) {
	var count int64
	if err := a.DB.Raw("SELECT COUNT(*) FROM users WHERE phone = $1 or email = $2", username, username).Scan(&count).Error; err != nil {
		return true, err
	}

	return count > 0, nil

}

func (a *authenticationRepository) GetUserDetailsByUsername(username string) (domain.User, error) {
	var model domain.User
	if err := a.DB.Raw("SELECT * FROM users WHERE phone = $1 or email = $2", username, username).Scan(&model).Error; err != nil {
		return domain.User{}, err
	}

	return model, nil

}
