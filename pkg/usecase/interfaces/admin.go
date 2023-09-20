package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type AdminUsecase interface {
	AdminLogin(context.Context, models.AdminLogin) (models.Tokens, error)
}
