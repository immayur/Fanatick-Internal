package firebase

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// Client is a Firebase client.
type Client struct {
}

// Client should be a token authenticator.
var _ fanatick.Authenticator = &Client{}

// Authenticate authenticates a token.
func (c *Client) Authenticate(token string) (*fanatick.User, error) {
	// TODO: implement
	return nil, nil
}
