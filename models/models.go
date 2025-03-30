package models

type ReqMail struct {
	SecretKey string `json:"secret_key" validate:"required"`
}
