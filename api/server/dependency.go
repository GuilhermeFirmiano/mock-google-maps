package server

import (
	"github.com/GuilhermeFirmiano/mock-google-maps/api/controllers"
	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/services"
	"github.com/sarulabs/di"
)

const (
	geocodeControllerDef = "geocode-controller"

	geocodeServiceDef = "geocode-service"
)

//RegisterDependencies ...
func (server *Server) RegisterDependencies() {

	server.AddDependency(di.Def{
		Name:  geocodeServiceDef,
		Scope: di.App,
		Build: func(container di.Container) (interface{}, error) {
			return services.NewGeocodeSerivce(), nil
		},
	})

	server.AddController(di.Def{
		Name:  geocodeControllerDef,
		Scope: di.App,
		Build: func(container di.Container) (interface{}, error) {
			geocodeService := container.Get(geocodeServiceDef).(services.GeocodeSerivce)
			return controllers.NewGeocodeController(geocodeService), nil
		},
	})

}
