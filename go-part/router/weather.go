package router

import (
	"github.com/gin-gonic/gin"
	"go-weather-api/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/weather", controller.GetWeather)
	return r
}
