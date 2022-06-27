package routes

import (
	"gin-go-api/controllers"

	"github.com/gin-gonic/gin"
)

func ObjectRoute(router *gin.Engine) {
	router.GET("/objects", controllers.GetObjects())
}
