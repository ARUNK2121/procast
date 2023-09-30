package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type WorkRepository interface {
	GetLeadByServiceAndLocation(service, location int) (int, error)
	GetDetailsOfAWork(int) (models.MinWorkDetails, error)
	GetImagesOfAWork(int) ([]string, error)
}
