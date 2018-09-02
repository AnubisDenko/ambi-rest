package main

type token struct{
	refreshToken string `json:"refresh_token"`
	scope string `json:"scope"`
	accessToken string `json:"access_token"`
	expiresIn int `json:"expires_in"`
	tokenType string `json:"token_type"`
}

type login struct{
	Username string `json:"username"`
	Password string `json:"password"`
}