package server

import (
	"github.com/GuilhermeFirmiano/mock-google-maps/api/controllers"
	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/services"
	"github.com/sarulabs/di"
)

const (
	geocodeControllerDef        = "geocode-controller"
	distanceMatrixControllerDef = "distanceMatrix-controller"

	geocodeServiceDef        = "geocode-service"
	distanceMatrixServiceDef = "distanceMatrix-service"
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
			service := container.Get(geocodeServiceDef).(services.GeocodeSerivce)
			return controllers.NewGeocodeController(service), nil
		},
	})

	server.AddDependency(di.Def{
		Name:  distanceMatrixServiceDef,
		Scope: di.App,
		Build: func(container di.Container) (interface{}, error) {
			return services.NewDistanceMatrixService(), nil
		},
	})

	server.AddController(di.Def{
		Name:  distanceMatrixControllerDef,
		Scope: di.App,
		Build: func(container di.Container) (interface{}, error) {
			service := container.Get(distanceMatrixServiceDef).(services.DistanceMatrixService)
			return controllers.NewDistanceMatrixController(service), nil
		},
	})

}
