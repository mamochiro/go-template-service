package container

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-template/internals/config"
	"go-template/internals/controller"
	"go-template/internals/infrastructure/database"
	grpcServer "go-template/internals/infrastructure/grpcServer"
	"go-template/internals/infrastructure/httpServer"
	"go-template/internals/infrastructure/jaeger"
	"go-template/internals/infrastructure/logrus"
	"go-template/internals/repository/postgres"
	servicePing "go-template/internals/service/ping"
	"go-template/internals/utils"
	"net/http"
)

var InfrastructureConstructor = []interface{}{
	grpcServer.NewServer,
	database.NewServerBase,
	http.NewServeMux,
	httpServer.NewServer,
	runtime.NewServeMux,
	jaeger.NewJaeger,
}

var ControllerConstructor = []interface{}{
	controller.NewHealthZController,
	controller.NewPingPongController,
}

var ServicesConstructors = []interface{}{
	servicePing.NewService,
}

var RepositoryConstructors = []interface{}{
	postgres.NewRepository,
}

var UtilConstructors = []interface{}{
	config.NewConfiguration,
	utils.NewUtils,
	utils.NewCustomValidator,
	logrus.NewLog,
}
