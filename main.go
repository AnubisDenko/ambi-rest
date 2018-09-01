package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", SayHello)
	router.Run(":" + port)
}

func SayHello(c *gin.Context){
	c.String(http.StatusOK, string("Test"))
}


