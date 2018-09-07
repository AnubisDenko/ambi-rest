package main

type token struct{
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type login struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type ambi struct{
	Location string `json:"location_name"`
	Room string `json:"room_name"`
	Value string `json:"value"`
}