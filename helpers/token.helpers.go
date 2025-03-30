package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func RefreshAccessToken() (*oauth2.Token, error) {

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:4001/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.readonly"},
	}

	token, err := loadToken("token.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load token: %v", err)
	}

	tokenSource := config.TokenSource(context.Background(), token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %v", err)
	}

	saveToken("token.json", newToken)
	return newToken, nil
}

func loadToken(filename string) (*oauth2.Token, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}

func saveToken(filename string, token *oauth2.Token) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(token)
}
