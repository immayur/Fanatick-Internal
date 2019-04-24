package fanatick

// TokenAuthenticator defines the operation to authenticate a token.
type TokenAuthenticator interface {
	Authenticate(token string) (*User, error)
}

// TokenGenerator defines the operation to generate an auth token.
type TokenGenerator interface {
	Generate(user *User) (string, error)
}
