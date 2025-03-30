package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"www.github.com/NirajSalunke/server/helpers"
)

var OauthConfig *oauth2.Config
var Emails []gin.H

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		helpers.PrintRed("Failed to Load Env")
		helpers.PrintRed(err.Error())
		return
	}
	helpers.PrintGreen("Env Loaded Successfully")

}

func Oauth() {
	credentials, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Error reading credentials.json: %v", err)
	}

	OauthConfig, err = google.ConfigFromJSON(credentials, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Error parsing credentials.json: %v", err)
	}
}
