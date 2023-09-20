package interfaces

import "github.com/ARUNK2121/procast/pkg/domain"

type AdminRepository interface {
	GetAdminDetailsByEmail(email string) (domain.Admin, error)
}
