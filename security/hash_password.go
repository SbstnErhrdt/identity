package security

import (
	"crypto/rand"
	"crypto/sha1"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"unicode"
)

// MinPasswordLength is the minimum length of a password.
const MinPasswordLength = 8

// ErrLength is an error that is returned when the password is not 8 characters long
var ErrLength = errors.New("password should have more than 8 characters.")

// errLength is an error that is returned when the password is not the length of the MinPasswordLength
func errLength() error {
	// adapt the error message based on the length of the password
	ErrLength = errors.New("password should have more than " + strconv.Itoa(MinPasswordLength) + " characters.")
	return ErrLength
}

// HashPassword hashes the users password and also generates a specific salt per user
func HashPassword(pepper string, password string, salt []byte) (hashedPassword []byte, hashedSalt []byte) {
	// Get random salt
	if len(salt) == 0 {
		salt = make([]byte, 64)
		if _, err := rand.Reader.Read(salt); err != nil {
			panic("random reader failed")
		}
	}
	hashedSalt = salt
	password = pepper + password
	hashedPassword = pbkdf2.Key([]byte(""+password), hashedSalt, 262144, 64, sha1.New)
	return
}

// ErrNumeric is an error that is returned when the password does not contain any numeric characters
var ErrNumeric = errors.New("password should have a minimum of 1 numeric character! e.g. 1,..,9,0.")

// ErrUppercase is an error that is returned when the password does not contain any uppercase characters
var ErrUppercase = errors.New("password should have a minimum of 1 upper case letter [A-Z].")

// ErrLowercase is an error that is returned when the password does not contain any lowercase characters
var ErrLowercase = errors.New("password should have a minimum of 1 lower case letter [a-z].")

// ErrSpecial is an error that is returned when the password does not contain any special characters
var ErrSpecial = errors.New("password should have a minimum of 1 special character: ~!@#$%^&*()-_+={}[]|:<>,./?...")

// CheckPasswordComplexity checks if the user password matches the conditions
func CheckPasswordComplexity(password string) (err error) {
	pwLength, number, upper, lower, special := ValidatePassword(password)
	if !pwLength {
		err = errLength()
		return err
	}
	if !number {
		err = ErrNumeric
		return err
	}
	if !upper {
		err = ErrUppercase
		return err
	}
	if !lower {
		err = ErrLowercase

		return err
	}
	if !special {
		err = ErrSpecial
		return err
	}
	return
}

// ValidatePassword validates the user's password
// seven or more characters, special, ...
func ValidatePassword(password string) (length, number, upper, lower, special bool) {
	letters := 0
	for _, s := range string(password) {
		switch {
		case unicode.IsNumber(s):
			number = true
			letters++
		case unicode.IsUpper(s):
			upper = true
			letters++
		case unicode.IsLower(s):
			lower = true
			letters++
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			special = true
			letters++
		case unicode.IsLetter(s) || s == ' ':
			letters++
		default:
			letters++
			//return false, false, false, false
		}
	}
	length = letters >= MinPasswordLength
	return
}
