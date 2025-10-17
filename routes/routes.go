package routes

import (
	"github.com/FallenCap/url-shortner/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/shorten", controller.ShortenUrl)

	router.GET("/:short", controller.RedirectToOriginal)

	return router
}
