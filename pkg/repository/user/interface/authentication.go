package interfaces

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type AuthenticationRepository interface {
	CheckIfPhoneNumberAlreadyExists(string) (bool, error)
	UserSignup(models.UserSignup) error
	CheckIfUserExistsByUsername(string) (bool, error)
	GetUserDetailsByUsername(string) (domain.User, error)
}
