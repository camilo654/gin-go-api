package main

import (
	"gin-go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.ObjectRoute(router)

	router.Run()
}
