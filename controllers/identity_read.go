package controllers

import (
	"github.com/SbstnErhrdt/identity/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

// ErrEmailIsAlreadyRegistered is returned when the email is already registered
var ErrEmailIsAlreadyRegistered = errors.New("email is already registered")

// ErrNoUserFound is returned when no user is found
var ErrNoUserFound = errors.New("no user has been found with this email")

// ErrUsernameLongerThan is returned when the username is shorter than 3 characters
var ErrUsernameLongerThan = errors.New("The username must be longer than 3 chars")

// ErrInvalidUsername is returned when the username is invalid
var ErrInvalidUsername = errors.New("The username must be alphanumeric and longer than 3 chars")

var regexUsername = regexp.MustCompile("^([0-9_A-Za-zÄÖÜäöüß]){3,60}$")

// CheckIfUsernameIsValid check if the username is valid
func CheckIfUsernameIsValid(userName string) (err error) {
	if len(userName) < 3 {
		return ErrUsernameLongerThan
	}
	if !regexUsername.Match([]byte(userName)) {
		return ErrInvalidUsername
	}
	return
}

// CheckIfEmailIsFree checks if the username is still free
func CheckIfEmailIsFree(service IdentityService, email string) (isFree bool, err error) {
	email = strings.ToLower(email)
	err = service.GetSQLClient().Limit(1).
		Where("lower(email) = lower(?)", email).
		Where("deleted_at is NULL"). // only users that are not deleted
		First(&models.Identity{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// username is free
		return true, nil
	} else if err != nil {
		log.Error(err)
		return false, ErrEmailIsAlreadyRegistered
	}
	return false, nil
}

// CheckUserLogin check if the credentials of a user are correct
func CheckUserLogin(service IdentityService, email string, password string) (result bool, err error) {
	user, err := GetIdentityByEmail(service, email)
	// Verify the password
	if VerifyPassword(service, user, password) {
		return true, nil
	} else {
		return false, errors.New("please check your password")
	}
}

// GetIdentityByEmail retrieves the identity from the database
func GetIdentityByEmail(service IdentityService, username string) (result *models.Identity, err error) {
	// Load identity from database
	identity := models.Identity{}
	err = service.GetSQLClient().Limit(1).
		Where("email = ?", username).
		Where("deleted_at is NULL"). // only users that are not deleted
		First(&identity).Error
	if err != nil {
		log.Error(err)
		return nil, ErrNoUserFound
	}
	return &identity, nil
}

// GetIdentityByUID retrieves the identity from the database
func GetIdentityByUID(service IdentityService, uid uuid.UUID) (result *models.Identity, err error) {
	// Load identity from database
	identity := models.Identity{}
	err = service.GetSQLClient().Limit(1).
		Where("uid = ?", uid.String()).
		Where("deleted_at is NULL"). // only users that are not deleted
		First(&identity).Error
	if err != nil {
		log.Error(err)
		return nil, ErrNoUserFound
	}
	return &identity, nil
}
