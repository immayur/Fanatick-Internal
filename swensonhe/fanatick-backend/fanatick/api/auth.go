package api

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
)

func AuthenticateTokenRequest(id string) (*fanatick.CommonTokenResponse, error) {
	events := fanatick.CommonTokenResponse{}
	events.Token = "2w1khd"
	return &events, nil
}
