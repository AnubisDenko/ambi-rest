package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

var jar,_ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
var client = &http.Client{Jar: jar, Timeout: time.Second * 10,}
var clientId = os.Getenv("AMBI_CLIENT_ID")
var clientSecret = os.Getenv("AMBI_CLIENT_SECRET")

const callBackUrl = "https://ambi-rest.herokuapp.com/secret"
const  authorizationRequestUrl = "https://api.ambiclimate.com/oauth2/authorize"
const  tokenRequestUrl = "https://api.ambiclimate.com/oauth2/token"

func StartAmbiAuthentication(ctx *gin.Context) {
	log.Println("Starting ambi authentication.")
	log.Println("Reading username and password from login request")
	log.Println("clientId", clientId, "client_secret", clientSecret)

	var credentials login
	if err:= ctx.ShouldBindJSON(&credentials); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No username password provided:" + err.Error()})
	}
	if "" == credentials.Username ||  credentials.Password == ""{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No credentials provided"})
		return
	}

	log.Println("logging into Ambi")
	ctx.JSON(http.StatusOK, gin.H{"status": "provided good info"})

	LoginAmbiServer(credentials)
	log.Println("Sending Request to receive authorization token")
	SendAuthorizationRequest()
}

func LoginAmbiServer(credentials login){
	const ambiLoginUrl = "https://api.ambiclimate.com/login"
	_, err := client.PostForm(ambiLoginUrl,url.Values{"email": {credentials.Username}, "password":{credentials.Password}} )
	if err != nil {
		log.Fatal("Failed to authentication with Ambi Climate for username", credentials.Username)
	}
}

func SendAuthorizationRequest(){
	resp, err := client.PostForm(authorizationRequestUrl, url.Values{
		"client_id":{clientId},
		"scope":{"email device_read ac_control"},
		"response_type":{"code"},
		"redirect_uri": {},
		"confirm": {"yes"}})

	if err != nil {
		log.Println("Received Error when requesting authorization token", err.Error())
	}
	PrintBody("SendAuthorizationRequest", resp.Body)
}

func AuthorizationTokenCallback(ctx *gin.Context){
	ctx.String(http.StatusOK, string("OK"))
	errorCode :=ctx.Query("error")
	if "" == errorCode {

	}
	code := ctx.Query("code")
	log.Println("Received Code", code)
	RequestAccessToken(code)
}

func RequestAccessToken(authorizationToken string){

	requestUrl ,_ := url.Parse(tokenRequestUrl)
	queryUrl := requestUrl.Query()
	queryUrl.Add("client_id",clientId)
	queryUrl.Add("redirect_uri",callBackUrl)
	queryUrl.Add("code",authorizationToken)
	queryUrl.Add("client_secret",clientSecret)
	queryUrl.Add("grant_type","authorization_code")
	requestUrl.RawQuery = queryUrl.Encode()

	log.Println(requestUrl.String())

	resp, err := client.Get(requestUrl.String())
	if err != nil {
		log.Println("Received Error when requesting authorization token", err.Error())
	}

	temp := GetBody("Body", resp.Body)
	log.Println("JSON String", temp)

	var myToken token
	err = json.Unmarshal([]byte(temp), &myToken)
	if err != nil {
		log.Fatal("Error while parsing access token",err)
	}
	log.Println("Access Token",myToken.AccessToken)
}