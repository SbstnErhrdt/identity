package identity_interface_graphql

import (
	"errors"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

const (
	// UserAgentContextKey is the key for the user agent in the context
	UserAgentContextKey = "USER_AGENT"
	// UserUidContextKey is the key for the user id in the context
	UserUidContextKey = "USER_UID"
	// UserIpContextKey is the key for the user ip in the context
	UserIpContextKey = "USER_IP"
	// OriginContextKey is the key for origin in the context
	OriginContextKey = "ORIGIN"
)

// ErrUserIdFromContext is returned when the user id is not found in the context
var ErrUserIdFromContext = errors.New("can not get USER_UID from context")

// GetUserUIDFromContext returns the user id from the context
func GetUserUIDFromContext(p *graphql.ResolveParams) (uid uuid.UUID, err error) {
	uid, ok := p.Context.Value(UserUidContextKey).(uuid.UUID)
	if !ok {
		err = ErrUserIdFromContext
		log.Error(err)
		return
	}
	return
}

// ErrAgentFromContext is returned when the user agent is not found in the context
var ErrAgentFromContext = errors.New("can not extract user agent from context")

// GetUserAgentFromContext returns the user agent from the context
func GetUserAgentFromContext(p *graphql.ResolveParams) (userAgent string, err error) {
	userAgent, ok := p.Context.Value(UserAgentContextKey).(string)
	if !ok {
		err = ErrAgentFromContext
		log.Error(err)
		return
	}
	return
}

// ErrIpFromContext is returned when the user ip is not found in the context
var ErrIpFromContext = errors.New("can not extract ip from context")

// GetIpFromContext returns the user ip from the context
func GetIpFromContext(p *graphql.ResolveParams) (ip string, err error) {
	ip, ok := p.Context.Value(UserIpContextKey).(string)
	if !ok {
		err = ErrIpFromContext
		log.Error(err)
		return
	}
	return
}

// ErrOriginFromContext is returned when the origin is not found in the context
var ErrOriginFromContext = errors.New("can not extract origin from context")

// GetOriginFromContext returns the origin from the context
func GetOriginFromContext(p *graphql.ResolveParams) (origin string, err error) {
	origin, ok := p.Context.Value(OriginContextKey).(string)
	if !ok {
		err = ErrOriginFromContext
		return
	}
	return
}

// ErrInvalidUid is returned when the uid is invalid
var ErrInvalidUid = errors.New("can not parse uid")

// ParseUIDFromArgs parses the uid from the context
func ParseUIDFromArgs(p *graphql.ResolveParams, key string) (uid uuid.UUID, err error) {
	uidString, ok := p.Args[key].(string)
	if !ok {
		err = ErrInvalidUid
		return
	}
	uid, err = uuid.Parse(uidString)
	return
}
