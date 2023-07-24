package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

// ErrTokenNameEmpty is thrown when the token name is empty
var ErrTokenNameEmpty = errors.New("token name is empty")

// ErrTokenExpirationDateInPast is thrown when the token expiration date is in the past
var ErrTokenExpirationDateInPast = errors.New("token expiration date is in the past")

// CreateApiToken creates a new api token for an identity
func CreateApiToken(service IdentityService, identityUID uuid.UUID, tokenName string, utcTokenExpirationDate time.Time) (token identity_models.IdentityApiToken, err error) {
	// init logger
	logger := log.WithFields(log.Fields{
		"identityUID": identityUID.String(),
		"process":     "CreateApiToken",
	})
	// check if the token name is empty
	if tokenName == "" {
		err = ErrTokenNameEmpty
		logger.WithError(err).Error("could not create api token with empty name")
		return
	}
	// check if the timestamp is in the past
	if utcTokenExpirationDate.Before(time.Now().UTC()) {
		err = ErrTokenExpirationDateInPast
		logger.WithError(err).Error("could not create api token in the past")
		return
	}
	// set expiration date to utc
	utcTokenExpirationDate = utcTokenExpirationDate.UTC()
	// generate token
	tokenString, tokenUID, err := security.GenerateJWTTokenWithExpirationData(identityUID, tokenName, map[string]interface{}{}, utcTokenExpirationDate)
	// create token db record
	token = identity_models.IdentityApiToken{
		IdentityUID:    identityUID,
		TokenUID:       tokenUID,
		Name:           tokenName,
		Token:          tokenString,
		ExpirationDate: utcTokenExpirationDate,
	}
	// create user
	logger.Debug("create api token in the database")
	err = service.GetSQLClient().Create(&token).Error
	if err != nil {
		logger.Error(err)
		return
	}

	return token, nil
}

// DeleteApiToken deletes an api token from the database
func DeleteApiToken(service IdentityService, identityUID uuid.UUID, tokenUID uuid.UUID) (err error) {
	// delete token
	err = service.GetSQLClient().
		Where("identity_uid = ?", identityUID.String()).
		Where("uid = ?", tokenUID.String()).
		Where("deleted_at is NULL").
		Delete(&identity_models.IdentityApiToken{}).Error
	if err != nil {
		service.GetLogger().
			WithError(err).
			WithField("identity_uid", identityUID).
			Error("could delete api tokens with uid")
		return err
	}
	return
}
