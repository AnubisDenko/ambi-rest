package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

var accessToken token
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/login", StartAmbiAuthentication)
	router.POST("/powerOff", PowerOff)
	router.GET("/secret", AuthorizationTokenCallback)
	router.Run(":" + port)
}

func GetBody(source string, body io.ReadCloser) string {
	buf := new (bytes.Buffer)
	buf.ReadFrom(body)
	bodyAsString :=buf.String()
	log.Println(source, bodyAsString)
	return bodyAsString
}

func PowerOff(ctx *gin.Context){
	log.Println("Calling PowerOff")
	log.Println("Have accessToken", accessToken.AccessToken)

	if accessToken.AccessToken == ""{
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Please login before issuing commands"})
		return
	}
	powerOffUrl := "https://api.ambiclimate.com/api/v1/device/power/off"
	var ambiDevice ambi
	if err:= ctx.ShouldBindJSON(&ambiDevice); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't translate request into JSON"})
		return
	}

	log.Println("Powering Off", ambiDevice)
	_, err := SendRequest(powerOffUrl, &ambiDevice)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Received Error from Ambi Server" + err.Error()})
	}
	ctx.JSON(http.StatusOK, nil)
}
