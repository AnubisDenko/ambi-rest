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
	router.GET("/secret", ReceiveSecret)
	router.Run(":" + port)
}

func SayHello(c *gin.Context){
	c.String(http.StatusOK, string("Test"))
}

func ReceiveSecret(c *gin.Context){
	c.String(http.StatusOK, string("OK"))
	log.Fatal(c)
}


