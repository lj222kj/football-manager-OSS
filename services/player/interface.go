package player

import "context"

type Service interface {
	Get(context.Context) (*Player, error)
	GetById() error
}
