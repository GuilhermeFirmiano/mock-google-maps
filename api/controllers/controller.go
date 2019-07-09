package controllers

import (
	"github.com/gin-gonic/gin"
)

//Controller ...
type Controller interface {
	RegisterRoutes(router *gin.RouterGroup)
}
