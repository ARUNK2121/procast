package providerRepository

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
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

func (p *profileRepository) CheckIfServiceIsAlreadyRegistered(user_id, service_id int) (bool, error) {
	var count int64
	if err := p.DB.Raw("SELECT COUNT(*) FROM probooks WHERE pro_id = $1 AND profession_id = $2", user_id, service_id).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (p *profileRepository) CheckIfDistrictIsAlreadyAdded(id, district int) (bool, error) {
	var count int64
	if err := p.DB.Raw("SELECT COUNT(*) FROM preferred_locations WHERE pro_id = $1 AND district_id = $2", id, district).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (p *profileRepository) AddService(user_id, service_id int) error {

	if err := p.DB.Exec("INSERT INTO probooks(pro_id,profession_id)VALUES($1,$2)", user_id, service_id).Error; err != nil {
		return err
	}

	return nil
}

func (p *profileRepository) AddLocation(id, district int) error {

	if err := p.DB.Exec("INSERT INTO preferred_locations(pro_id,district_id)VALUES($1,$2)", id, district).Error; err != nil {
		return err
	}

	return nil
}

func (p *profileRepository) DeleteService(user_id, service_id int) error {

	if err := p.DB.Exec("DELETE FROM probooks WHERE pro_id = $1 AND profession_id = $2", user_id, service_id).Error; err != nil {
		return err
	}

	return nil
}

func (p *profileRepository) RemovePreferredLocation(id, district int) error {

	if err := p.DB.Exec("DELETE FROM preferred_locations WHERE pro_id = $1 AND district_id = $2", id, district).Error; err != nil {
		return err
	}

	return nil
}

func (p *profileRepository) GetAllServiceIdsOfAProvider(id int) ([]int, error) {

	var services []int

	if err := p.DB.Raw("SELECT profession_id FROM probooks WHERE pro_id = $1", id).Scan(&services).Error; err != nil {
		return []int{}, err
	}

	return services, nil
}

func (p *profileRepository) FindServiceDetailsFromID(id int) (domain.Profession, error) {

	var service domain.Profession

	if err := p.DB.Raw("SELECT * FROM professions WHERE id = $1", id).Scan(&service).Error; err != nil {
		return domain.Profession{}, err
	}

	return service, nil
}

func (p *profileRepository) GetAllPreferredLocations(id int) ([]int, error) {

	var locations []int

	if err := p.DB.Raw("SELECT district_id FROM preferred_locations WHERE pro_id = $1", id).Scan(&locations).Error; err != nil {
		return []int{}, err
	}

	return locations, nil
}

func (p *profileRepository) GetLocationDetails(id int) (models.GetLocations, error) {

	var service models.GetLocations

	if err := p.DB.Raw("select district,states.state AS state from districts join states on states.id=districts.state_id where districts.id = $1", id).Scan(&service).Error; err != nil {
		return models.GetLocations{}, err
	}

	return service, nil
}

func (p *profileRepository) FindProviderDetails(id int) (domain.Provider, error) {

	var pro domain.Provider
	if err := p.DB.Raw("SELECT * FROM providers WHERE id = $1", id).Scan(&pro).Error; err != nil {
		return domain.Provider{}, err
	}

	return pro, nil
}

func (p *profileRepository) GetRatingsOfAllRecordsOfAProvider(id int) ([]int, error) {

	var ratings []int

	if err := p.DB.Raw(`SELECT ratings.rating FROM works JOIN ratings ON works.id = ratings.work_id JOIN providers ON providers.id = works.pro_id WHERE providers.id = $1`, id).Scan(&ratings).Error; err != nil {
		return []int{}, err
	}

	return ratings, nil
}
