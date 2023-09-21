package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
)

type AdminRepository interface {
	GetAdminDetailsByEmail(context.Context, string) (domain.Admin, error)
	CreateNewAdmin(context.Context, domain.Admin) error
	CountOfAdminByEmail(context.Context, string) (int64, error)
	DeleteAdmin(context.Context, int) error
}
