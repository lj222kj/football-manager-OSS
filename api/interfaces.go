package api

import (
	"context"

	"github.com/lj222kj/services/player"
)

type PlayerService interface {
	Get(ctx context.Context) (*player.Player, error)
}
