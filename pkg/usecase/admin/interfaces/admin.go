package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type AdminUsecase interface {
	AdminLogin(context.Context, models.AdminLogin) (models.DoubleTokens, error)
	CreateNewAdmin(context.Context, domain.Admin) error
	DeleteAdmin(context.Context, int) error
}
