package providerRepository

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
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
	if err := a.DB.Raw("SELECT COUNT(*) FROM providers WHERE phone = ?", phone).Scan(&count).Error; err != nil {
		return true, err
	}

	return count > 0, nil

}

func (a *authenticationRepository) Register(model models.ProviderRegister) (int, error) {
	err := a.DB.Exec("INSERT INTO providers(name,email,password,phone)VALUES($1,$2,$3,$4)", model.Name, model.Email, model.Password, model.Phone).Error
	if err != nil {
		return 0, err
	}

	var id int64
	err = a.DB.Raw("Select id from providers where phone = $1", model.Phone).Scan(&id).Error
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (a *authenticationRepository) UploadDocument(id int, file string) error {
	if err := a.DB.Exec("INSERT INTO id_proofs(pro_id,id_proof) VALUES($1,$2)", id, file).Error; err != nil {
		return err
	}

	return nil
}

func (a *authenticationRepository) CheckIfProviderExistsByUsername(username string) (bool, error) {
	var count int64
	if err := a.DB.Raw("SELECT COUNT(*) FROM providers WHERE phone = $1 or email = $2", username, username).Scan(&count).Error; err != nil {
		return true, err
	}

	return count > 0, nil

}

func (a *authenticationRepository) GetProviderDetailsByUsername(username string) (domain.Provider, error) {
	var model domain.Provider
	if err := a.DB.Raw("SELECT * FROM providers WHERE phone = $1 or email = $2", username, username).Scan(&model).Error; err != nil {
		return domain.Provider{}, err
	}

	return model, nil

}
