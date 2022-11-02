package grpcserver

import (
	"go-template/internals/config"
	"go-template/internals/controller"
	apiV1 "go-template/pkg/api/v1"
	grpcHealthV1 "go-template/pkg/grpc/health/v1"

	"go-template/internals/utils"
	"google.golang.org/grpc"
)

type Server struct {
	Config       config.Configuration
	Server       *grpc.Server
	HealthCtrl   *controller.HealthZController
	PingPongCtrl *controller.PingPongController
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	validator *utils.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor()

	s := &Server{
		Server:       grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:       config,
		HealthCtrl:   healthCtrl,
		PingPongCtrl: pingPongCtrl,
	}

	s.Configure()

	return s
}
