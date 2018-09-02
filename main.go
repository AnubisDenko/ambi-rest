package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/login", StartAmbiAuthentication)
	router.GET("/secret", AuthorizationTokenCallback)
	router.GET("/accessToken", ReceiveAccessToken)
	router.Run(":" + port)
}

func PrintBody(body io.ReadCloser){
	buf := new (bytes.Buffer)
	buf.ReadFrom(body)
	bodyAsString :=buf.String()

	log.Println(bodyAsString)
}
