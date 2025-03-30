package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"www.github.com/NirajSalunke/server/config"
)

func loadToken(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func HandleLogin(ctx *gin.Context) {
	authURL := config.OauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusFound, authURL)
}

func HandleCallBack(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	token, err := config.OauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed"})
		return
	}

	file, err := os.Create("token.json")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token file"})
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(token)

	ctx.String(http.StatusOK, "Login successful! You can close this window.")
}
