package controllers

import "github.com/gin-gonic/gin"

//StaticMapsController ...
type StaticMapsController struct {
}

//NewStaticMapsController ...
func NewStaticMapsController() *StaticMapsController {
	return &StaticMapsController{}
}

// RegisterRoutes ...
func (ctrl *StaticMapsController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("staticmap", ctrl.get)
}

func (ctrl *StaticMapsController) get(c *gin.Context) {
	c.File("pkg/assets/staticmap.png")
}
