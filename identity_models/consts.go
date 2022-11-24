package identity_models

type TokenType string

const (
	LoginToken    TokenType = "LOGIN"
	PasswordReset TokenType = "PW_RESET"
)
