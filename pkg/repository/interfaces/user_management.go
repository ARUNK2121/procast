package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type UserManagementRepository interface {
	GetProviders(context.Context) ([]domain.Provider, error)
	MakeProviderVerified(ctx context.Context, id int) error
	RevokeVerification(ctx context.Context, id int) error
	GetUsers(context.Context) ([]models.UserDetails, error)
}
