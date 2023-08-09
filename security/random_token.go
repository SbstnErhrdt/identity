package security

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log/slog"
	"strings"
)

// ErrCannotGenerateRandomBytes is the error returned when we cannot generate a random string
var ErrCannotGenerateRandomBytes = errors.New("cannot generate random bytes")

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		slog.With("err", err).Error("cannot generate random bytes")
		err = ErrCannotGenerateRandomBytes
		return nil, err
	}
	return b, nil
}

// ErrCannotGenerateRandomString is the error returned when we cannot generate a random string
var ErrCannotGenerateRandomString = errors.New("cannot generate random token")

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	if err != nil {
		slog.With("err", err).Error("cannot generate random string")
		return "", ErrCannotGenerateRandomString
	}
	// outlook has problems with == in the base64 string
	// replace = with _ to avoid this
	token := strings.Replace(base64.URLEncoding.EncodeToString(b), "=", "_", -1)
	return token, err
}
