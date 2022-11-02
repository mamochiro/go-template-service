package container

import (
	"fmt"
	"net/http"

	"go-template/internals/config"
	"go-template/internals/controller"
	"go-template/internals/infrastructure/database"
	grpcServer "go-template/internals/infrastructure/grpcServer"
	httpServer "go-template/internals/infrastructure/httpServer"
	"go-template/internals/infrastructure/jaeger"
	"go-template/internals/infrastructure/logrus"
	"go-template/internals/repository/postgres"
	servicePing "go-template/internals/service/ping"
	"go-template/internals/service/ping/wrapper"
	"go-template/internals/utils"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {
	if err := c.container.Provide(wrapper.WrapInfo, dig.Name("wrapperInfo")); err != nil {
		return err
	}
	servicesConstructors := []interface{}{
		config.NewConfiguration,
		grpcServer.NewServer,
		database.NewServerBase,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,
		logrus.NewLog,
		controller.NewHealthZController,
		controller.NewPingPongController,
		servicePing.NewService,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	return nil
}

func (c *Container) Start() error {
	fmt.Println("Start Container")

	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *utils.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		fmt.Printf("%s", err)

		return err
	}

	return nil
}

// MigrateDB ...
func (c *Container) MigrateDB() error {
	fmt.Println("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
