package services

import (
	"github.com/stretchr/testify/assert"
	"net/mail"
	"testing"
)

func TestSend(t *testing.T) {
	ass := assert.New(t)
	err := SendEmail(mail.Address{
		Name:    "test@erhardt.net",
		Address: "test@erhardt.net",
	}, mail.Address{
		Name:    "erhardt.sebastian@gmail.com",
		Address: "erhardt.sebastian@gmail.com",
	}, "Hello World Test", "Wow it works")
	ass.NoError(err)
}
