package controller

import (
	"fmt"
	"time"

	"github.com/francomatheus/casa-do-codigo-golang/model"
	"github.com/francomatheus/casa-do-codigo-golang/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AutorController(client *mongo.Client, c *gin.Context) {
	var autorRequest model.AutorRequest
	validationError := c.ShouldBindJSON(&autorRequest)
	if validationError != nil {
		c.AbortWithStatus(400)
	}

	var autorDocument repository.AutorEntity = repository.AutorEntity{
		Nome:      autorRequest.Nome,
		Email:     autorRequest.Email,
		Descricao: autorRequest.Descricao,
		Instante:  time.Now(),
	}

	saveError := autorDocument.Save(client)
	if saveError != nil {
		fmt.Println(saveError)
	}

	c.JSON(201, autorRequest)
}
