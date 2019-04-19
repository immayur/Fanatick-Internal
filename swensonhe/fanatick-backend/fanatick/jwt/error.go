package jwt

// Error represents an error.
type Error string

// The errors.
const (
	ErrorUnauthorized            = Error("unauthorized")
	ErrorUnexpectedSigningMethod = Error("unexpected_signing_method")
	ErrorUnexpecteClaimsType     = Error("unexpected_claims_type")
	ErrorTokenStringUnsigned     = Error("token_string_unsigned")
)

// Error implements the error interface
func (e Error) Error() string {
	return string(e)
}
