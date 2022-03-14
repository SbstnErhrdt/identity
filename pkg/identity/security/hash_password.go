package security

import (
	"crypto/rand"
	"crypto/sha1"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
	"unicode"
)

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

// CheckPasswordComplexity checks if the user password matches the conditions
func CheckPasswordComplexity(password string) (err error) {
	sevenMore, number, upper, lower, special := ValidatePassword(password)
	if !sevenMore {
		err = errors.New("password should have more than 7 characters.")
		return err
	}
	if !number {
		err = errors.New("password should have a minimum of 1 numeric character! e.g. 1,..,9,0.")
		return err
	}
	if !upper {
		err = errors.New("password should have a minimum of 1 upper case letter [A-Z].")
		return err
	}
	if !lower {
		err = errors.New("password should have a minimum of 1 lower case letter [a-z].")
		return err
	}
	if !special {
		err = errors.New("password should have a minimum of 1 special character: ~!@#$%^&*()-_+={}[]|:<>,./?...")
		return err
	}
	return
}

// ValidatePassword validates the user's password
// seven or more characters, special, ...
func ValidatePassword(password string) (sevenOrMore, number, upper, lower, special bool) {
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
	sevenOrMore = letters >= 7
	return
}
