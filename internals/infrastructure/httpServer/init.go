package httpServer

import (
	"context"
	"go-template/internals/controller"
	"net/http"
	"strconv"

	"go-template/internals/config"
	apiV1 "go-template/pkg/api/v1"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	Config   config.Configuration
	Server   *runtime.ServeMux
	HttpMux  *http.ServeMux
	PingPong *controller.PingPongController
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
	apiV1.RegisterPingPongServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux, pingPong *controller.PingPongController) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:   config,
		Server:   rmux,
		HttpMux:  httpMux,
		PingPong: pingPong,
	}
	s.Configure(context.Background(), opts)
	return s
}
