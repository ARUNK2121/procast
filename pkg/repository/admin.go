package repository

import (
	"context"
	"errors"

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
	if ctx.Err() != nil {
		return domain.Admin{}, errors.New("timeout")
	}

	return model, nil
}

func (a *AdminRepository) CreateNewAdmin(ctx context.Context, model domain.Admin) error {
	tx := a.DB.Begin()
	//insert new admin to admin table
	err := tx.Exec("INSERT INTO admins (name,email,password,previlege) VALUES($1,$2,$3,$4)", model.Name, model.Email, model.Password, model.Previlege).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (a *AdminRepository) CountOfAdminByEmail(ctx context.Context, email string) (int64, error) {
	var count int64
	err := a.DB.Raw("SELECT COUNT(*) FROM admins WHERE email = $1", email).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (a *AdminRepository) DeleteAdmin(ctx context.Context, id int) error {
	tx := a.DB.Begin()
	var count int
	err := tx.Exec("DELETE FROM admins WHERE id = $1", id).Scan(&count).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if ctx.Err() != nil {
		tx.Rollback()
		return errors.New("request timeout")
	}

	tx.Commit()
	return nil
}
