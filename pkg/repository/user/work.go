package user_repository

import (
	"fmt"

	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
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

	err := w.DB.Exec("INSERT INTO works(street,district_id,state_id,target_profession_id,user_id)VALUES($1,$2,$3,$4,$5)", model.Street, model.DistrictID, model.StateID, model.TargetProfessionID, model.UserID).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *workRepository) GetAllWorksOfAUser(id int) ([]int, error) {
	var works []int
	if err := p.DB.Raw(`SELECT id FROM works WHERE user_id = $1`, id).Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (p *workRepository) FindUsername(id int) (string, error) {
	var name string
	if err := p.DB.Raw(`SELECT name FROM users WHERE id = $1`, id).Scan(&name).Error; err != nil {
		return "", err
	}

	return name, nil
}

func (p *workRepository) GetDetailsOfAWork(id int) (models.MinWorkDetails, error) {
	var model models.MinWorkDetails
	if err := p.DB.Raw(`SELECT works.id,works.street,districts.district as district,states.state as state,professions.profession as profession,users.name AS user,works.work_status 
	FROM works 
	JOIN districts ON districts.id=works.district_id 
	JOIN states ON states.id=works.state_id 
	JOIN professions ON professions.id=works.target_profession_id 
	JOIN users ON users.id=works.user_id 
	WHERE works.id=$1`, id).Scan(&model).Error; err != nil {
		return models.MinWorkDetails{}, err
	}

	fmt.Println("model", model)

	return model, nil
}

func (p *workRepository) GetImagesOfAWork(id int) ([]string, error) {
	var images []string
	if err := p.DB.Raw("SELECT image FROM workspace_images WHERE work_id = $1", id).Scan(&images).Error; err != nil {
		return []string{}, err
	}

	return images, nil
}

func (p *workRepository) FindProviderIdFromWork(id int) (int, error) {
	var pro_id int
	if err := p.DB.Raw("SELECT pro_id FROM works WHERE id = $1", id).Scan(&pro_id).Error; err != nil {
		return 0, nil
	}

	return pro_id, nil
}

func (p *workRepository) FindProviderName(pro_id int) (string, error) {
	var name string
	if err := p.DB.Raw("SELECT name FROM providers WHERE id = $1", pro_id).Scan(&name).Error; err != nil {
		return "", err
	}

	return name, nil
}
