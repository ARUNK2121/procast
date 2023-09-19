package handler

import (
	"github.com/ARUNK2121/procast/pkg/usecase/interfaces"
)

type AdminHandler struct {
	Usecase interfaces.AdminUsecase
}

func NewAdminHandler(usecase interfaces.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		Usecase: usecase,
	}
}
