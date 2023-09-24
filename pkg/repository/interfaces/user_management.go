package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
)

type UserManagementRepository interface {
	GetProviders(context.Context) ([]domain.Provider, error)
	MakeProviderVerified(ctx context.Context, id int) error
	RevokeVerification(ctx context.Context, id int) error
}
