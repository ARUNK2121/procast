package providerhandler

import interfaces "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"

type AuthenticationHandler struct {
	Usecase interfaces.AuthenticationUsecase
}

func NewAuthenticationHandler(use interfaces.AuthenticationUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{
		Usecase: use,
	}
}
