package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)
//const tokenRequestUrl = "https://ambi-rest.herokuapp.com/"
const tokenRequestUrl = "https://api.ambiclimate.com/oauth2/authorize?client_id=cHKV&redirect_uri=https%3A%2F%2Fambi-rest.herokuapp.com%2Fsecret&response_type=code"
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", SayHello)
	router.GET("/secret", ReceiveSecret)
	//RequestToken()
	router.Run(":" + port)


}

func RequestToken(){

	timer:= time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		log.Println("Sending Token Request")
		resp, err := http.Get(tokenRequestUrl)
		if  err != nil {
			log.Fatal("Error",err)
		}

		if resp == nil {
			log.Fatal("No response")
		}
		ReadBody(resp.Body)
	}()
}

func ReadBody(body io.ReadCloser){
	buf := make([]byte, 1024)
	length,_  := body.Read(buf)
	bodyAsString := string(buf[0:length])
	log.Println("Body",bodyAsString)
}


func SayHello(c *gin.Context){
	c.String(http.StatusOK, string("Test"))
}

func ReceiveSecret(c *gin.Context){
	c.String(http.StatusOK, string("OK"))
	code := c.Query("code")
	log.Println(code)
}


