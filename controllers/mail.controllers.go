package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"www.github.com/NirajSalunke/server/config"
	"www.github.com/NirajSalunke/server/helpers"
	"www.github.com/NirajSalunke/server/models"
)

var IsNewMailInData bool

func getHeader(headers []*gmail.MessagePartHeader, name string) string {
	for _, h := range headers {
		if h.Name == name {
			return h.Value
		}
	}
	return "(No Subject)"
}
func getMessageBody(parts []*gmail.MessagePart) string {
	var body string
	for _, part := range parts {
		if part.MimeType == "text/plain" {
			data, err := base64.URLEncoding.DecodeString(part.Body.Data)
			if err != nil {
				continue
			}
			body = string(data)
			break
		} else if part.Parts != nil {
			body = getMessageBody(part.Parts)
			if body != "" {
				break
			}
		}
	}
	return body
}

func GetAllMails(ctx *gin.Context) {
	token, err := loadToken("token.json")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found. Please login first."})
		return
	}
	if token.Expiry.Before(time.Now()) {
		fmt.Println("Access token expired. Refreshing...")
		token, err = helpers.RefreshAccessToken()
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to refresh token. Please re-authenticate."})
			return
		}
	}
	client := config.OauthConfig.Client(context.Background(), token)
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve Gmail client"})
		return
	}

	user := "me"
	msgList, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve messages"})
		return
	}

	existingEmailIDs := make(map[string]bool)
	for _, email := range config.Emails {
		if id, ok := email["id"].(string); ok {
			existingEmailIDs[id] = true
		}
	}

	IsNewMailInData = false
	for _, msg := range msgList.Messages {
		if _, exists := existingEmailIDs[msg.Id]; exists {
			continue
		}

		message, err := srv.Users.Messages.Get(user, msg.Id).Format("full").Do()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to retrieve message details for ID %s", msg.Id)})
			return
		}

		from := getHeader(message.Payload.Headers, "From")
		to := getHeader(message.Payload.Headers, "To")
		subject := getHeader(message.Payload.Headers, "Subject")
		date := getHeader(message.Payload.Headers, "Date")
		body := getMessageBody(message.Payload.Parts)

		config.Emails = append(config.Emails, gin.H{
			"id":      msg.Id,
			"from":    from,
			"to":      to,
			"subject": subject,
			"date":    date,
			"body":    body,
		})
		existingEmailIDs[msg.Id] = true
		IsNewMailInData = true
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "update": IsNewMailInData})
}

func GetMail(ctx *gin.Context) {

	var req models.ReqMail
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if req.SecretKey != os.Getenv("SECRET_KEY") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid secret key"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"mails":   config.Emails,
	})

}
