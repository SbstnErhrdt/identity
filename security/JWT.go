package security

import (
	"github.com/SbstnErhrdt/env"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

// ParseToken takes a string and extracts the token
func ParseToken(tokenString string) (claims map[string]interface{}, err error) {
	claims = map[string]interface{}{}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			return claims, err
		}
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.New("unexpected signing method: " + token.Header["alg"].(string))
			log.Error(err)
			return claims, err
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(env.FallbackEnvVariable("SECURITY_JWT_SECRET", "SECRET")), nil
	})
	// check token
	if token == nil {
		err = errors.New("no token present")
		log.Error(err)
		return
	} else {
		if !token.Valid {
			err = errors.New("token not valid")
			log.Error(err)
			return
		} else {
			claims = token.Claims.(jwt.MapClaims)
			return
		}
	}
}

// GenerateJWTToken generate token for as user and a payload
func GenerateJWTToken(subjectUID uuid.UUID, audience string, payload map[string]interface{}) (tokenString string, tokenUID uuid.UUID, err error) {
	tokenUID = uuid.New()
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     subjectUID.String(),
		"iss":     env.FallbackEnvVariable("SECURITY_JWT_ISSUER", "ERHARDT"),
		"aud":     audience,
		"exp":     time.Now().UTC().Unix() + int64(60*60*24*7), // 7 days
		"iat":     time.Now().UTC().Unix(),
		"nbf":     time.Now().UTC().Unix(),
		"jti":     tokenUID.String(),
		"payload": payload,
	})
	// build string
	tokenString, err = token.SignedString([]byte(env.FallbackEnvVariable("SECURITY_JWT_SECRET", "SECRET")))
	// Sign and get the complete encoded token as a string using the secret
	return
}

// GeneratePasswordResetToken generate a password reset token with an expiration date
func GeneratePasswordResetToken(email string, expirationDate time.Time) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": env.FallbackEnvVariable("SECURITY_JWT_ISSUER", "ERHARDT"),
		"aud": "PASSWORD_RESET",
		"exp": expirationDate.UTC().Unix(),
		"iat": time.Now().UTC().Unix(),
		"nbf": time.Now().UTC().Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(env.FallbackEnvVariable("SECURITY_JWT_SECRET", "SECRET"))
}

// GenerateEmailChangeToken generate a token to verify a new email address
func GenerateEmailChangeToken(email string, expirationDate time.Time) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": env.FallbackEnvVariable("SECURITY_JWT_ISSUER", "ERHARDT"),
		"aud": "PASSWORD_RESET",
		"exp": expirationDate.Unix(),
		"iat": time.Now().UTC().Unix(),
		"nbf": time.Now().UTC().Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(env.FallbackEnvVariable("SECURITY_JWT_SECRET", "SECRET"))
}
