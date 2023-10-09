package interfaces

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type ProfileRepository interface {
	CheckIfServiceIsAlreadyRegistered(user_id, service_id int) (bool, error)
	AddService(user_id, service_id int) error
	DeleteService(user_id, service_id int) error
	CheckIfDistrictIsAlreadyAdded(int, int) (bool, error)
	AddLocation(int, int) error
	RemovePreferredLocation(id, district int) error
	GetAllServiceIdsOfAProvider(id int) ([]int, error)
	FindServiceDetailsFromID(int) (domain.Profession, error)
	GetAllPreferredLocations(int) ([]int, error)
	GetLocationDetails(int) (models.GetLocations, error)
	FindProviderDetails(id int) (domain.Provider, error)
	GetRatingsOfAllRecordsOfAProvider(id int) ([]int, error)
}
