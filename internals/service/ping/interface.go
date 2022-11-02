package ping

import (
	"context"

	"go-template/internals/model"
)

//go:generate mockery --name=Service
type Service interface {
	Ping(ctx context.Context, input model.Ping) (*model.Pong, error)
}
