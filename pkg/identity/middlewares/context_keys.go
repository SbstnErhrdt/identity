package middlewares

type ContextKey string

const (
	UserAgent ContextKey = "USER_AGENT"
	UserIP    ContextKey = "USER_IP"
)
