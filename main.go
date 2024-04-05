package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// "url-shortener/repositories"
	"url-shortener/routes"
	"url-shortener/utils"
)

func main() {
	router := SetupAppRouter()
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func SetupAppRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	apiRoute := router.Group("api")
	shortenRoute := router.Group("")

	routes.InitApiRoutes(apiRoute)
	routes.InitShortenRoutes(shortenRoute)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    "NOT_FOUND",
			"status":  http.StatusNotFound,
			"message": "The requested resource was not found.",
		})

		defer ctx.AbortWithStatus(http.StatusNotFound)
	})

	return router
}
