package ping

import (
	"go-template/internals/repository/postgres"
)

type PongService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &PongService{
		repository: r,
	}
}
