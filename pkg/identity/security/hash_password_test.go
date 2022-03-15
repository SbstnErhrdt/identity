package security

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestHashPassword(t *testing.T) {
	ass := assert.New(t)
	pw := "Hello World"
	hashPw, hashSalt := HashPassword("aa", pw, []byte{})
	hashPw2, hashSalt := HashPassword("aa", pw, hashSalt)
	ass.True(reflect.DeepEqual(hashPw, hashPw2))
}

func TestValidatePassword(t *testing.T) {

	ass := assert.New(t)

	Password := []byte("MaggieThatcherIs110%Sexy!")
	seven, number, upper, lower, special := ValidatePassword(string(Password))
	ass.True(seven)
	ass.True(number)
	ass.True(upper)
	ass.True(lower)
	ass.True(special)

	Password = []byte("password")
	seven, number, upper, lower, special = ValidatePassword(string(Password))
	ass.True(seven)
	ass.False(number)
	ass.False(upper)
	ass.True(lower)
	ass.False(special)

	Password = []byte("zagrageNkot7xdx!")
	seven, number, upper, lower, special = ValidatePassword(string(Password))
	ass.True(seven)
	ass.True(number)
	ass.True(upper)
	ass.True(lower)
	ass.True(special)

	Password = []byte("mF^4fM)P!")
	seven, number, upper, lower, special = ValidatePassword(string(Password))
	ass.True(seven)
	ass.True(number)
	ass.True(upper)
	ass.True(lower)
	ass.True(special)

}
