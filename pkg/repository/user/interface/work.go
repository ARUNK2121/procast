package interfaces

import "github.com/ARUNK2121/procast/pkg/domain"

type WorkRepository interface {
	ListNewOpening(domain.Work) error
}
