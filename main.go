package main

import (
	"news/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.Get)
	router.POST("/", routes.Post)

	router.Run()
}
