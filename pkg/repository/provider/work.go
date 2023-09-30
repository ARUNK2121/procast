package providerRepository

import (
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
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

func (p *workRepository) GetLeadByServiceAndLocation(service, location int) (int, error) {
	var id int64
	if err := p.DB.Raw("SELECT id FROM works WHERE target_profession_id = $1 AND district = $2 AND work_status = $3", service, location, "listed").Scan(&id).Error; err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *workRepository) GetDetailsOfAWork(id int) (models.MinWorkDetails, error) {
	var model models.MinWorkDetails
	if err := p.DB.Raw(`SELECT works.id,works.street,districts.district,states.state,professions.profession,users.name AS user,works.work_status 
	FROM works 
	JOIN districts ON districts.id=works.district_id 
	JOIN states ON states.id=works.state_id 
	JOIN professions ON professions.id=works.target_profession_id 
	JOIN users ON users.id=works.user_id 
	WHERE works.id=$1`, id).Scan(&model).Error; err != nil {
		return models.MinWorkDetails{}, err
	}

	return model, nil
}

func (p *workRepository) GetImagesOfAWork(id int) ([]string, error) {
	var images []string
	if err := p.DB.Raw("SELECT image FROM workspace_images WHERE work_id = $1", id).Scan(&images).Error; err != nil {
		return []string{}, err
	}

	return images, nil
}

func (p *workRepository) PlaceBid(model models.PlaceBid) error {
	if err := p.DB.Exec(`INSERT INTO bids(work_id,pro_id,estimate,description) VALUES($1,$2,$3,$4)`, model.WorkID, model.ProID, model.Estimate, model.Description).Error; err != nil {
		return err
	}

	return nil
}

func (p *workRepository) ReplaceBidWithNewBid(model models.PlaceBid) error {
	if err := p.DB.Exec(`UPDATE bids SET estimate = $1, description = $2 WHERE pro_id = $3 AND work_id = $4`, model.Estimate, model.Description, model.ProID, model.WorkID).Error; err != nil {
		return err
	}

	return nil
}