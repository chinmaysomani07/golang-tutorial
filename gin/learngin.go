package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type identity struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"identity": identity{Name: "Chinmay", Age: 23, Gender: "Male"},
		})
	})
	server.Run(":8080")
}
