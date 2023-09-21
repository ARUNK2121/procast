package repository

import (
	"context"

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

func (a *AdminRepository) GetAdminDetailsByEmail(ctx context.Context, email string) (domain.Admin, error) {
	var model domain.Admin
	err := a.DB.Raw("SELECT * FROM admins WHERE email = $1", email).Scan(&model).Error
	if err != nil {
		return domain.Admin{}, err
	}

	return model, nil
}

func (a *AdminRepository) CreateNewAdmin(ctx context.Context, model domain.Admin) error {
	//insert new admin to admin table
	err := a.DB.Exec("INSERT INTO admins (id,name,email,password,previlege) VALUES($1,$2,$3,$4,$5)", model.ID, model.Name, model.Email, model.Password, model.Previlege).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *AdminRepository) CountOfAdminByEmail(email string) (int, error) {
	var count int
	err := a.DB.Raw("SELECT COUNT(*) FROM admins WHERE email = $1", email).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
