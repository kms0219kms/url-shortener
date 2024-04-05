package routes

import (
	"net/http"
	shortenService "url-shortener/services"

	"github.com/gin-gonic/gin"
)

func InitShortenRoutes(route *gin.RouterGroup) {
	shortenService := shortenService.NewService()

	route.GET("/:shortened", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, shortenService.RedirectToOriginal(ctx.Params.ByName("shortened")))
	})
}
