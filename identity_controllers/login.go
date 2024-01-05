package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"gorm.io/gorm"
	"time"
)

// ErrExternalLoginFailedInvalidUserOrPassword is returned when the user or password is invalid
var ErrExternalLoginFailedInvalidUserOrPassword = errors.New("login failed; Invalid user ID or password")

// ErrExternalLoginFailed is returned when the login failed and the reason should not be exposed to the user
var ErrExternalLoginFailed = errors.New("login failed; please try again")

// ErrExternalUserBlocked is returned when the user is blocked
var ErrExternalUserBlocked = errors.New("this identity can not login. Please contact the support")

// ErrExternalUserCleared is returned when the user is not cleared
var ErrExternalUserCleared = errors.New("this identity can not login at the moment. It must be cleared by the admin first")

// ErrEmailNotVerified is returned when the user email is not verified
var ErrEmailNotVerified = errors.New("the email address is not confirmed yet")

// ErrIdentityNotFound is returned when the identity is not found
var ErrIdentityNotFound = errors.New("identity not found")

// ErrWrongPassword is returned when the password is wrong
var ErrWrongPassword = errors.New("wrong password")

// ErrExternalUserLoginNotPossible is returned when the user can not log in
var ErrExternalUserLoginNotPossible = errors.New("login currently not possible. Please try again later")

// ErrLoginTokenExpirationInPast is returned when the login token expiration date is in the past
var ErrLoginTokenExpirationInPast = errors.New("login token expiration date is in the past")

// Login logs in a user and returns a JWT token
func Login(service IdentityService, emailAddress, password, userAgent, ip string) (token string, err error) {
	// sanitize input
	emailAddress = SanitizeEmail(emailAddress)
	// init logger
	logger := service.GetLogger().With(
		"identity", emailAddress,
		"process", "Login",
	)
	// check login data
	if len(emailAddress) == 0 {
		err = ErrNoEmail
		logger.With("err", err).Error("no email address")
		return "", err
	}
	if len(password) == 0 {
		err = ErrNoPassword
		logger.With("err", err).Error("no password")
		return "", err
	}
	// track login attempt
	logger.Debug("track login attempt")
	loginAttempt := identity_models.IdentityLogin{
		Email:     emailAddress,
		UserAgent: userAgent,
		IP:        ip,
	}
	err = service.GetSQLClient().Create(&loginAttempt).Error
	if err != nil {
		logger.With("err", err).Error("could not track login attempt")
		return "", ErrExternalUserLoginNotPossible
	}
	// Check if identity exists
	// init identity object
	identity := &identity_models.Identity{}

	// Find the identity in the database
	logger.Debug("find identity")
	err = service.GetSQLClient().
		Where("email = ?", emailAddress).
		First(&identity).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrIdentityNotFound
		logger.With("err", err).Error("could not find identity")
		// OWASP: generic error message if user does not exist
		err = ErrExternalLoginFailedInvalidUserOrPassword
		return "", err
	} else if err != nil {
		logger.With("err", err).Error("could not find identity")
		// OWASP: generic error message
		err = ErrExternalLoginFailed
		return "", err
	}

	// check if identity is not blocked
	logger.Debug("check if identity is not blocked")
	if identity.Blocked {
		err = ErrExternalUserBlocked
		logger.With("err", err).Warn("identity is blocked")
		return "", err
	}

	// check if identity is already not cleared
	logger.Debug("check if identity is cleared")
	if !identity.Cleared {
		err = ErrExternalUserCleared
		logger.With("err", err).Warn("identity is not cleared")
		return "", err
	}

	// check if identity is active
	logger.Debug("check if identity is activated")
	if !identity.Active {
		err = ErrEmailNotVerified
		logger.With("err", err).Warn("identity is not activated")
		return "", err
	}

	// verify the password
	logger.Debug("verify password")
	if !(VerifyPassword(service, identity, password)) {
		err = ErrWrongPassword
		logger.With("err", err).With("securityIncident", 1).Error("wrong password")
		// OWASP: return generic error message
		err = ErrExternalLoginFailedInvalidUserOrPassword
		return "", err
	}

	// generate token
	logger.Debug("generate token")
	audience := env.FallbackEnvVariable("SECURITY_JWT_AUDIENCE", "APP")
	loginTokenExpiration := service.GetExpirationLoginDuration()
	// check if login token expiration is greater than 0
	if loginTokenExpiration <= 0 {
		err = ErrLoginTokenExpirationInPast
		logger.With("err", err).Error("login token expiration date is in the past")
		return "", err
	}
	// add the duration to the current time
	utcLoginTokenExpirationDate := time.Now().UTC().Add(loginTokenExpiration)
	token, tokenUID, errToken := security.GenerateJWTTokenWithExpirationDate(identity.UID, audience, map[string]interface{}{}, utcLoginTokenExpirationDate)
	if errToken != nil {
		logger.With("err", errToken).Error("could not generate token")
		// OWASP: generic error message
		err = ErrExternalLoginFailed
		return "", err
	}

	// save token id in token table
	tokenMeta := identity_models.IdentityTokenMeta{
		TokenUID:  tokenUID,
		TokenType: identity_models.LoginToken,
	}
	errTokenMeta := service.GetSQLClient().Create(&tokenMeta).Error
	if errTokenMeta != nil {
		logger.With("err", errTokenMeta).Error("could not generate token meta entry")
		// OWASP: generic error message
		err = ErrExternalLoginFailed
		return "", err
	}

	// update login attempt
	logger.Debug("update login attempt")
	loginAttempt.IdentityUID = &identity.UID
	service.GetSQLClient().Save(&loginAttempt)
	// return the token
	return token, nil
}
