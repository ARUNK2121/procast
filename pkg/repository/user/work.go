package user_repository

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	"gorm.io/gorm"
)

type workRepository struct {
	DB *gorm.DB
}

func NewWorkRepository(db *gorm.DB) interfaces.WorkRepository {
	return &workRepository{
		DB: db,
	}
}

func (w *workRepository) ListNewOpening(model domain.Work) error {

	err := w.DB.Exec("INSERT INTO works(street,district_id,state_id,target_profession_id,user_id)VALUES($1,$2,$3,$4)", model.Street, model.District, model.State, model.TargetProfessionID, model.UserID).Error
	if err != nil {
		return err
	}

	return nil
}
