package main

import (
	"log"
	"net/http"
	"net/url"
)

func SendRequest(urlString string, device *ambi) (*http.Response, error){
	requestUrl, _ := url.Parse(urlString)
	queryUrl:= requestUrl.Query()
	queryUrl.Add("room_name", device.Room)
	queryUrl.Add("location_name", device.Location)
	queryUrl.Add("multiple", "False")
	requestUrl.RawQuery = queryUrl.Encode()

	log.Println("Starting to send request to", requestUrl.String())

	req, _ := http.NewRequest("GET",requestUrl.String(),nil)
	req.Header.Set("Accept","application/json")
	req.Header.Set("Authorization","Bearer " + accessToken.AccessToken)
	return client.Do(req)
}