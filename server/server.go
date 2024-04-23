package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	log.Println("Starting the server at")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
