package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	"time"
)

// ErrTokenNameEmpty is thrown when the token name is empty
var ErrTokenNameEmpty = errors.New("token name is empty")

// ErrTokenExpirationDateInPast is thrown when the token expiration date is in the past
var ErrTokenExpirationDateInPast = errors.New("token expiration date is in the past")

// CreateApiToken creates a new api token for an identity
func CreateApiToken(service IdentityService, identityUID uuid.UUID, tokenName string, utcTokenExpirationDate time.Time) (token *identity_models.IdentityApiToken, err error) {
	// init logger
	logger := service.GetLogger().With(
		"identityUID", identityUID.String(),
		"process", "CreateApiToken",
	)
	// check if the token name is empty
	if tokenName == "" {
		err = ErrTokenNameEmpty
		logger.With("err", err).Error("could not create api token with empty name")
		return
	}
	// check if the timestamp is in the past
	if utcTokenExpirationDate.Before(time.Now().UTC()) {
		err = ErrTokenExpirationDateInPast
		logger.With("err", err).Error("could not create api token in the past")
		return
	}
	// set expiration date to utc
	utcTokenExpirationDate = utcTokenExpirationDate.UTC()
	audience := env.FallbackEnvVariable("SECURITY_JWT_API_AUDIENCE", "API")
	// generate token
	tokenString, tokenUID, err := security.GenerateJWTTokenWithExpirationData(identityUID, audience, map[string]interface{}{
		"tokenName": tokenName,
	}, utcTokenExpirationDate)
	// create token db record
	t := identity_models.IdentityApiToken{
		IdentityUID:    identityUID,
		TokenUID:       tokenUID,
		Name:           tokenName,
		Token:          tokenString,
		ExpirationDate: utcTokenExpirationDate,
	}
	// create user
	logger.Debug("create api token in the database")
	err = service.GetSQLClient().Create(&t).Error
	if err != nil {
		logger.With("err", err).Error("could not create api token in the database")
		return
	}

	return &t, nil
}

// DeleteApiToken deletes an api token from the database
func DeleteApiToken(service IdentityService, identityUID uuid.UUID, tokenUID uuid.UUID) (err error) {
	// init logger
	logger := service.GetLogger().With(
		"identityUID", identityUID.String(),
		"tokenUID", tokenUID.String(),
		"process", "DeleteApiToken",
	)
	// delete token
	err = service.GetSQLClient().
		Where("identity_uid = ?", identityUID.String()).
		Where("uid = ?", tokenUID.String()).
		Where("deleted_at is NULL").
		Delete(&identity_models.IdentityApiToken{}).Error
	if err != nil {
		logger.
			With("err", err).
			With("identity_uid", identityUID).
			Error("could delete api tokens with uid")
		return err
	}
	return
}
