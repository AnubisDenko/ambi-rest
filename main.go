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
	router.POST("/feedback",AdjustComfort)
	router.POST("/comfort",ComfortMode)
	router.POST("/temperatureLower",TemperatureLower)
	router.GET("/secret", AuthorizationTokenCallback)
	router.Run(":" + port)
}

func GetBody(source string, body io.ReadCloser) string {
	buf := new (bytes.Buffer)
	buf.ReadFrom(body)
	bodyAsString :=buf.String()
	log.Println(source, bodyAsString)
	body.Close()
	return bodyAsString
}

func TemperatureLower(ctx *gin.Context){
	log.Println("Calling Temperature Lower")
	comfort := "https://api.ambiclimate.com/api/v1/device/mode/away_temperature_lower"
	HandlePost(ctx, comfort)
}

func ComfortMode(ctx *gin.Context){
	log.Println("Calling ComfortMode")
	comfort := "https://api.ambiclimate.com/api/v1/device/mode/comfort"
	HandlePost(ctx, comfort)
}

func AdjustComfort(ctx *gin.Context){
	log.Println("Calling AdjustComfort")
	comfort := "https://api.ambiclimate.com/api/v1/user/feedback"
	HandlePost(ctx, comfort)
}


func PowerOff(ctx *gin.Context){
	log.Println("Calling PowerOff")
	powerOffUrl := "https://api.ambiclimate.com/api/v1/device/power/off"
	HandlePost(ctx, powerOffUrl)
}

func HandlePost(ctx *gin.Context, ambiUrl string){
	var ambiDevice ambi
	if accessToken.AccessToken == "" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Please login before issuing commands"})
		return
	}

	if err:= ctx.ShouldBindJSON(&ambiDevice); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't translate request into JSON"})
	}
	_, err := SendRequest(ambiUrl, &ambiDevice)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, nil)
}