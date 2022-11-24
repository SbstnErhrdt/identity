package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ErrLoginFailedInvalidUserOrPassword is returned when the user or password is invalid
var ErrLoginFailedInvalidUserOrPassword = errors.New("login failed; Invalid user ID or password")

// ErrLoginFailed is returned when the login failed and the reason should not be exposed to the user
var ErrLoginFailed = errors.New("login failed; please try again")

// ErrUserBlocked is returned when the user is blocked
var ErrUserBlocked = errors.New("this identity can not login. please contact the support")

// ErrEmailNotVerified is returned when the user email is not verified
var ErrEmailNotVerified = errors.New("the email address is not confirmed yet")

// ErrIdentityNotFound is returned when the identity is not found
var ErrIdentityNotFound = errors.New("identity not found")

// ErrWrongPassword is returned when the password is wrong
var ErrWrongPassword = errors.New("wrong password")

// Login logs in a user and returns a JWT token
func Login(service IdentityService, emailAddress, password, userAgent, ip string) (token string, err error) {
	// sanitize input
	emailAddress = SanitizeEmail(emailAddress)
	// init logger
	logger := log.WithFields(log.Fields{
		"identity": emailAddress,
		"process":  "Login",
	})
	// check login data
	if len(emailAddress) == 0 {
		logger.Error(ErrNoEmail)
		return "", ErrNoEmail
	}
	if len(password) == 0 {
		logger.Error(ErrNoPassword)
		return "", ErrNoPassword
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
		logger.Error(err)
		return "", err
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
		logger.Error(err)
		// OWASP: generic error message if user does not exist
		err = ErrLoginFailedInvalidUserOrPassword
		return "", err
	} else if err != nil {
		logger.Error(err)
		// OWASP: generic error message
		err = ErrLoginFailed
		return "", err
	}

	// check if identity is not blocked
	logger.Debug("check if identity is not blocked")
	if identity.Blocked {
		err = ErrUserBlocked
		logger.Warn(err)
		return "", err
	}

	// check if identity is active
	logger.Debug("check if identity is activated")
	if !identity.Active {
		err = ErrEmailNotVerified
		logger.Warn(err)
		return "", err
	}

	// verify the password
	logger.Debug("verify password")
	if !(VerifyPassword(service, identity, password)) {
		err = ErrWrongPassword
		logger.Error(err)
		// OWASP: return generic error message
		err = ErrLoginFailedInvalidUserOrPassword
		return "", err
	}

	// generate token
	log.Debug("generate token")
	audience := env.FallbackEnvVariable("SECURITY_JWT_AUDIENCE", "API")
	token, tokenUID, errToken := security.GenerateJWTToken(identity.UID, audience, map[string]interface{}{})
	if errToken != nil {
		logger.Error(errToken)
		// OWASP: generic error message
		err = ErrLoginFailed
		return "", errToken
	}

	// save token id in token table
	tokenMeta := identity_models.IdentityTokenMeta{
		TokenUID:  tokenUID,
		TokenType: identity_models.LoginToken,
	}
	errTokenMet := service.GetSQLClient().Create(&tokenMeta).Error
	if errToken != nil {
		logger.Error(errTokenMet)
		// OWASP: generic error message
		err = ErrLoginFailed
		return "", errTokenMet
	}

	// update login attempt
	logger.Debug("update login attempt")
	loginAttempt.IdentityUID = identity.UID
	service.GetSQLClient().Save(&loginAttempt)
	// return the token
	return token, nil
}
