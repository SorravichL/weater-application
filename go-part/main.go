package main

import (
	"go-weather-api/router"
	"go-weather-api/utils"
)

func main() {
	utils.LoadEnv()
	r := router.SetupRouter()
	r.Run(":" + utils.GetEnv("PORT", "8080"))
}
