package firebase

import (
	"context"

	"fmt"

	"firebase.google.com/go/auth"
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// Client is a Firebase client.
type Client struct {
	*auth.Client
	Ctx context.Context
}

// Client should be a token authenticator.
//var _ fanatick.Authenticator = &Client{}

// Authenticate authenticates a token.
func (c *Client) Authenticate(token string) (*fanatick.User, error) {
	// TODO: implement
	return nil, nil
}

// check if user exists in firebase storage
func (c *Client) UserExists(firebaseUUID string) bool {
	_, err := c.GetUser(c.Ctx, firebaseUUID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// used to verify firebase token if needed
	//_, err := c.VerifyIDToken(c.Ctx, idToken)
	//if err != nil {
	//	return false
	//}

	return true
}
