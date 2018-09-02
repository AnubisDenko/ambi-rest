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
	router.GET("/secret", ReceivedAuthorizationToken)
	router.Run(":" + port)
}

func PrintBody(body io.ReadCloser){
	buf := new (bytes.Buffer)
	buf.ReadFrom(body)
	bodyAsString :=buf.String()

	log.Println(bodyAsString)
	//buf := make([]byte, 1024)
	//length,_  := body.Read(buf)
	//bodyAsString := string(buf[0:length])
	//log.Println("Body",bodyAsString)
}



// old stuff below
//func TimerExample(){
//	timer:= time.NewTimer(2 * time.Second)
//	go func() {
//		<-timer.C
//		log.Println("Sending Token Request")
//		resp, err := http.Get(tokenRequestUrl)
//		if  err != nil {
//			log.Fatal("Error",err)
//		}
//
//		if resp == nil {
//			log.Fatal("No response")
//		}
//		PrintBody(resp.Body)
//	}()
//}
//
//
//

//
//
//func SayHello(c *gin.Context){
//	c.String(http.StatusOK, string("Test"))
//}
//


