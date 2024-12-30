package player

import (
	"context"

	positions "github.com/lj222kj/xpkg/positions"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s service) Get(ctx context.Context) (*Player, error) {
	return &Player{
		ID:       "019418c1-e397-7f42-a08d-96fdb31c5b54",
		Name:     "Elnour Ibrahimović",
		Age:      43,
		Position: positions.Striker,
		Attributes: Attributes{
			Speed:   1.0,
			Stamina: 1.0,
			Passing: 1.0,
		},
	}, nil
}

func (s service) GetById() error {
	return nil
}
