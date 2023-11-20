package response

import "github.com/chicken-afk/go-fiber/pkg/models"

type LoginResponse struct {
	TokenType string      `json:"token_type"`
	Token     string      `json:"token"`
	User      models.User `json:"user"`
}
