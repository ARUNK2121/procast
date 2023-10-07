package userhandler

import (
	interfaces "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
)

type ProfileHandler struct {
	usecase interfaces.ProfileUsecase
}

func NewProfileHandler(use interfaces.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		usecase: use,
	}
}
