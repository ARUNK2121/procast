package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type UserManagementUsecase interface {
	GetProviders(context.Context) ([]models.ProviderDetails, error)
	MakeProviderVerified(ctx context.Context, id int) error
	RevokeVerification(ctx context.Context, id int) error
	GetUsers(context.Context) ([]models.UserDetails, error)
	BlockUser(ctx context.Context, id int) error
	UnBlockUser(ctx context.Context, id int) error
	GetAllPendingVerifications(context.Context) ([]models.Verification, error)
	ViewVerificationRequest(context.Context, int) (models.VerificationDetails, error)
}
