package identity_models

// TokenType is the type of the token
type TokenType string

const (
	LoginToken    TokenType = "LOGIN"
	ApiToken      TokenType = "API"
	PasswordReset TokenType = "PW_RESET"
)
