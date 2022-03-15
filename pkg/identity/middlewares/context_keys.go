package middlewares

// ContextKey defines a key type for context keys
type ContextKey string

const (
	// UserAgent is a key for user agent in context
	UserAgent ContextKey = "USER_AGENT"
	// UserIP is a key for user ip in context
	UserIP ContextKey = "USER_IP"
)
