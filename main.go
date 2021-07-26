package main

import (
	"github.com/francomatheus/casa-do-codigo-golang/configuration"
	"github.com/francomatheus/casa-do-codigo-golang/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	client := configuration.NewConnection()

	server := gin.Default()

	v1 := server.Group("/cdc/v1")
	v1.POST("/autor", func(c *gin.Context) {
		controller.AutorController(client, c)
	})

	server.Run(":8080")
}
