package controllers

import (
	"net/http"

	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/services"
	"github.com/gin-gonic/gin"
)

// DistanceMatrixController ...
type DistanceMatrixController struct {
	service services.DistanceMatrixService
}

// NewDistanceMatrixController ...
func NewDistanceMatrixController(service services.DistanceMatrixService) *DistanceMatrixController {
	return &DistanceMatrixController{
		service: service,
	}
}

// RegisterRoutes ...
func (ctrl *DistanceMatrixController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("distancematrix/json", ctrl.get)
}

func (ctrl *DistanceMatrixController) get(c *gin.Context) {
	distancematrix := ctrl.service.Get()
	c.JSON(http.StatusOK, distancematrix)
}
