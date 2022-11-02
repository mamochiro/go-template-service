package wrapper

import (
	service "go-template/internals/service/ping"

	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperInfo"`
	Service service.Service
}

func WrapInfo(service service.Service) service.Service {
	return &Wrapper{
		Service: service,
	}
}
