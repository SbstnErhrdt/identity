package identity_middlewares

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// ReadUserIP extract ip from http request
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// UserAgentAndIpMiddleware extracts the ip and the client
// from the request and adds it to the context
func UserAgentAndIpMiddleware() gin.HandlerFunc {
	logger := slog.With(
		"middleware", "identity_user_agent_and_ip",
	)
	return func(c *gin.Context) {
		// Extract ConfirmationIP
		logger.Debug("set ip")
		c.Set(string(UserIP), ReadUserIP(c.Request))
		// Extract ConfirmationUserAgent
		logger.Debug("set user agent")
		c.Set(string(UserAgent), c.Request.UserAgent())
		// next middlewares or handlers
		c.Next()
	}
}
