package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type ProfileUsecase interface {
	AddService(int, int) error
	DeleteService(int, int) error
	AddPreferredWorkingLocation(int, int) error
	RemovePreferredLocation(id, district int) error
	GetMyServices(int) ([]models.GetServices, error)
	GetAllPreferredLocations(int) ([]models.GetLocations, error)
}
