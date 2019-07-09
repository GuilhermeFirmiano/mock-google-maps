package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	ctrls "github.com/GuilhermeFirmiano/mock-google-maps/api/controllers"
	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/logging"
	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/settings"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag"
)

//Server wraps API configurations.
type Server struct {
	Engine   *gin.Engine
	Settings *settings.Settings

	DIBuilder *di.Builder
	Container di.Container

	router      *gin.RouterGroup
	controllers []string
}

var (
	containerKey       = "di-container"
	errContainerNotSet = errors.New("container not set in request scope")
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	swag.Register(swag.Name, NewSwaggerDoc())
}

//Configure creates a new API server
func Configure(settings *settings.Settings) *Server {
	server := &Server{}
	server.Settings = settings

	logging.LogWithApplication(server.Settings.ApplicationName)

	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	server.DIBuilder = builder

	server.Engine = gin.New()
	server.Engine.Use(gin.Recovery())

	server.Engine.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	server.router = server.Engine.Group("/maps/api/")

	server.router.Use(server.containerHandler())

	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return server
}

//AddDependency register a new dependency in DI container.
func (server *Server) AddDependency(def di.Def) error {
	return server.DIBuilder.Add(def)
}

//AddController register a new controller to be handled.
func (server *Server) AddController(def di.Def) error {
	server.controllers = append(server.controllers, def.Name)
	return server.DIBuilder.Add(def)
}

//Run starts the server.
func (server *Server) Run() {
	server.ensureContainer()

	srv := http.Server{
		Addr:    server.Settings.Host,
		Handler: server.Engine,
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		sig := <-sigs
		logging.LogInfo("caught sig: %+v", sig)
		logging.LogInfo("waiting 5 seconds to finish processing")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logging.LogWith(err).Error("shutdown error")
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.LogWith(err).Info("startup error")
	}
}

//RegisterControllers register all routes added.
func (server *Server) RegisterControllers() {
	server.ensureContainer()
	for _, ctrl := range server.extractControllers() {
		ctrl.RegisterRoutes(server.router)
	}
}

//Container return DI Container defined in request scope.
func Container(c *gin.Context) (di.Container, error) {
	container, ok := c.Get(containerKey)

	if !ok {
		return nil, errContainerNotSet
	}

	di, ok := container.(di.Container)

	if !ok {
		return nil, errContainerNotSet
	}

	return di, nil
}

func (server *Server) containerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if server.Container != nil {
			container, err := server.Container.SubContainer()

			if err != nil {
				panic(err)
			}

			defer container.Delete()
			c.Set(containerKey, container)
		}
		c.Next()
	}
}

func (server *Server) extractControllers() []ctrls.Controller {
	var c []ctrls.Controller

	for _, name := range server.controllers {
		def, err := server.Container.SafeGet(name)
		ctrl, ok := def.(ctrls.Controller)

		if err != nil {
			panic(err)
		}

		if !ok {
			panic("Defs added in AddController must implements Controller")
		}

		c = append(c, ctrl)
	}

	return c
}

func (server *Server) ensureContainer() {
	if server.Container == nil {
		server.Container = server.DIBuilder.Build()
	}
}
