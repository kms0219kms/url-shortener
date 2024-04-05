package routes

import (
	"net/http"
	"url-shortener/dtos"
	shortenService "url-shortener/services"

	"github.com/gin-gonic/gin"
)

func InitApiRoutes(route *gin.RouterGroup) {
	shortenService := shortenService.NewService()

	route.PUT("/v2/:shortened", func(ctx *gin.Context) {
		var input dtos.ShortenRequestDto
		ctx.ShouldBindJSON(&input)

		ctx.String(http.StatusOK, shortenService.CreateShorten(ctx.Params.ByName("shortened"), input.Original))
	})

	route.GET("/v2/:shortened", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, shortenService.GetShorten(ctx.Params.ByName("shortened")))
	})

	route.PATCH("/v2/:shortened", func(ctx *gin.Context) {
		var input dtos.ShortenRequestDto
		ctx.ShouldBindJSON(&input)

		ctx.String(http.StatusOK, shortenService.UpdateShorten(ctx.Params.ByName("shortened"), input.Original))
	})

	route.DELETE("/v2/:shortened", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, shortenService.DeleteShorten(ctx.Params.ByName("shortened")))
	})

	route.Any("/v1/*_", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"code": "NOT_ACCEPTABLE",
			"status": http.StatusNotAcceptable,
			"message": "This API version is deprecated. Please use the latest version.",
		})

		defer ctx.AbortWithStatus(http.StatusNotAcceptable)
	})
}
