package main

import (
	"os"
)

type Settings struct {
	AppDomain        string
	AuthDomain       string
	AuthAudience     string
	AuthIssuer       string
	AuthClientId     string
	AuthClientSecret string
}

func NewSettings() *Settings {
	return &Settings{
		AppDomain:        os.Getenv("APP_DOMAIN"),
		AuthDomain:       os.Getenv("AUTH_DOMAIN"),
		AuthAudience:     os.Getenv("AUTH_AUDIENCE"),
		AuthIssuer:       os.Getenv("AUTH_ISSUER"),
		AuthClientId:     os.Getenv("AUTH_CLIENT_ID"),
		AuthClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
	}
}
