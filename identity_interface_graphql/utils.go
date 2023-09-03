package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_config"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"log/slog"
)

// ErrUserIdFromContext is returned when the user id is not found in the context
var ErrUserIdFromContext = errors.New("can not get IDENTITY_UID from context")

// GetIdentityUIDFromContext returns the user id from the context
func GetIdentityUIDFromContext(p *graphql.ResolveParams) (uid uuid.UUID, err error) {
	uid, ok := p.Context.Value(identity_config.IdentityUIDContextKey).(uuid.UUID)
	if !ok {
		err = ErrUserIdFromContext
		slog.With("err", err).Error("can not get IDENTITY_UID from context")
		return
	}
	return
}

// ErrAgentFromContext is returned when the user agent is not found in the context
var ErrAgentFromContext = errors.New("can not extract user agent from context")

// GetUserAgentFromContext returns the user agent from the context
func GetUserAgentFromContext(p *graphql.ResolveParams) (userAgent string, err error) {
	userAgent, ok := p.Context.Value(identity_config.UserAgentContextKey).(string)
	if !ok {
		err = ErrAgentFromContext
		slog.With("err", err).Error("can not get agent from context")
		return
	}
	return
}

// ErrIpFromContext is returned when the user ip is not found in the context
var ErrIpFromContext = errors.New("can not extract ip from context")

// GetIpFromContext returns the user ip from the context
func GetIpFromContext(p *graphql.ResolveParams) (ip string, err error) {
	ip, ok := p.Context.Value(identity_config.UserIpContextKey).(string)
	if !ok {
		err = ErrIpFromContext
		slog.With("err", err).Error("can not get ip from context")
		return
	}
	return
}

// ErrOriginFromContext is returned when the origin is not found in the context
var ErrOriginFromContext = errors.New("can not extract origin from context")

// GetOriginFromContext returns the origin from the context
func GetOriginFromContext(p *graphql.ResolveParams) (origin string, err error) {
	origin, ok := p.Context.Value(identity_config.OriginContextKey).(string)
	if !ok {
		err = ErrOriginFromContext
		slog.With("err", err).Error("can not get origin from context")
		return
	}
	return
}

// ErrInvalidUid is returned when the uid is invalid
var ErrInvalidUid = errors.New("can not parse uid from arguments")

// ParseUIDFromArgs parses the uid from the context
func ParseUIDFromArgs(p *graphql.ResolveParams, key string) (uid uuid.UUID, err error) {
	uidString, ok := p.Args[key].(string)
	if !ok {
		err = ErrInvalidUid
		slog.With("err", err).Error("can not get uid from context")
		return
	}
	uid, err = uuid.Parse(uidString)
	return
}
