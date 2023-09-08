package identity_controllers

import (
	"github.com/pkg/errors"
	"regexp"
)

// ErrEmailIsAlreadyRegistered is returned when the email is already registered
var ErrEmailIsAlreadyRegistered = errors.New("email is already registered")

// ErrNoUserFound is returned when no user is found
var ErrNoUserFound = errors.New("no user has been found with this email")

// ErrUsernameLongerThan is returned when the username is shorter than 3 characters
var ErrUsernameLongerThan = errors.New("the username must be longer than 3 chars")

// ErrInvalidUsername is returned when the username is invalid
var ErrInvalidUsername = errors.New("the username must be alphanumeric and longer than 3 chars")

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

// ErrCheckPassword is returned when the password is wrong
var ErrCheckPassword = errors.New("please check your password")

// CheckUserLogin check if the credentials of a user are correct
func CheckUserLogin(service IdentityService, email string, password string) (result bool, err error) {
	email = SanitizeEmail(email)
	user, err := GetIdentityByEmail(service, email)
	if err != nil {
		service.GetLogger().With("err", err).With("email", email).Error("could not find user with email")
	}
	// Verify the password
	if VerifyPassword(service, user, password) {
		return true, nil
	} else {
		err = ErrCheckPassword
		return false, err
	}
}
