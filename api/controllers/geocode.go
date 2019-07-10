package controllers

import (
	"net/http"

	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/services"
	"github.com/gin-gonic/gin"
)

// GeocodeController ...
type GeocodeController struct {
	service services.GeocodeSerivce
}

// NewGeocodeController ...
func NewGeocodeController(service services.GeocodeSerivce) *GeocodeController {
	return &GeocodeController{
		service: service,
	}
}

// RegisterRoutes ...
func (ctrl *GeocodeController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("geocode/json", ctrl.get)
}

func (ctrl *GeocodeController) get(c *gin.Context) {
	geocode := ctrl.service.Get()
	c.JSON(http.StatusOK, geocode)
}
