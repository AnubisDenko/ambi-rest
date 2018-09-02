package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)
//var cookieUrl, _ = url.Parse("https://api.ambiclimate.com")
var  tokenRequestUrl = "https://api.ambiclimate.com/oauth2/authorize"

var jar,_ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
var c = &http.Client{Jar: jar}

func StartAmbiAuthentication(c *gin.Context) {
	//myCookieJar := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	log.Println("Starting ambi authentication.")
	log.Println("Reading username and password from login request")
	var credentials login
	if err:= c.ShouldBindJSON(&credentials); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "No username password provided:" + err.Error()})
	}
	if "" == credentials.Username ||  credentials.Password == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No credentials provided"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "provided good info"})
	log.Println("logging into Ambi")
	cookie:= LoginAmbiServer(credentials)
	log.Println("Received", cookie)
	log.Println("Sending Request to receive authorization token")
	SendAuthorizationRequest(cookie)
}

func SendAuthorizationRequest(cookie *http.Cookie){

	resp, err := c.PostForm(tokenRequestUrl, url.Values{
		"client_id":{"cHKV"},
		"scope":{"email device_read ac_control"},
		"response_type":{"code"},
		"redirect_uri": {"https://ambi-rest.herokuapp.com/secret"},
		"confirm": {"yes"}})

	if err != nil {
		log.Println("Received Error when requesting authorization token", err.Error())
	}
	PrintBody(resp.Body)
}

func LoginAmbiServer(credentials login) *http.Cookie{
	const ambiLoginUrl = "https://api.ambiclimate.com/login"
	resp, err := c.PostForm(ambiLoginUrl,url.Values{"email": {credentials.Username}, "password":{credentials.Password}} )

	if err != nil {
		log.Fatal("Failed to authentication with Ambi Climate for username", credentials.Username)
	}
	for _, element := range resp.Cookies(){
		if element.Name == "session"{
			//sessionCookie := element
			return element
		}
	}
	log.Fatal("Didn't find any session cookie so everything else won't work")
	return nil
}

func ReceivedAuthorizationToken(c *gin.Context){
	c.String(http.StatusOK, string("OK"))
	code := c.Query("code")
	log.Println("Received Code", code)
}

func RequestAuthorization(token string){

}


