package ping

import (
	"context"
	"go-template/internals/model"
)

func (p PongService) Ping(ctx context.Context, request model.Ping) (*model.Pong, error) {
	return &model.Pong{Success: true}, nil
}
