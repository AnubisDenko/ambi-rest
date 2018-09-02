package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)
var  tokenRequestUrl = "https://api.ambiclimate.com/oauth2/authorize"

var jar,_ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
var c = &http.Client{Jar: jar}

func StartAmbiAuthentication(c *gin.Context) {
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
	log.Println("Sending Request to receive authorization token")
	SendAuthorizationRequest()
}

func SendAuthorizationRequest(){

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

func AuthorizationTokenCallback(c *gin.Context){
	c.String(http.StatusOK, string("OK"))
	code := c.Query("code")
	log.Println("Received Code", code)
	RequestAccessToken(code)
}

func RequestAccessToken(authorizationToken string){
	resp, err := c.PostForm(tokenRequestUrl, url.Values{
		"client_id":{"cHKV"},
		"redirect_uri": {"https://ambi-rest.herokuapp.com/accessToken"},
		"code":{authorizationToken},
		"client_secret":{"9a9p4"},
		"grant_type": {"authorization_code"}})

	if err != nil {
		log.Println("Received Error when requesting authorization token", err.Error())
	}
	PrintBody(resp.Body)
}

func ReceiveAccessToken(c *gin.Context){
	var accessToken token
	if err:= c.ShouldBindJSON(&accessToken); err != nil{
		log.Fatal("Couldn't read access token")
	}

	log.Println(accessToken)
}

