package interfaces

import "github.com/ARUNK2121/procast/pkg/domain"

type WorkUsecase interface {
	ListNewOpening(domain.Work) error
}
