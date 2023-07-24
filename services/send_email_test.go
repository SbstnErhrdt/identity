package services

import (
	"github.com/SbstnErhrdt/env"
	"github.com/stretchr/testify/assert"
	"net/mail"
	"testing"
)

func init() {
	env.LoadEnvFiles("../test/test.env")
}

func TestSend(t *testing.T) {
	ass := assert.New(t)
	err := SendEmail(mail.Address{
		Name:    "test@erhardt.net",
		Address: "test@erhardt.net",
	}, mail.Address{
		Name:    "test@erhardt.net",
		Address: "test@erhardt.net",
	}, "Hello World Test", "Wow it works")
	ass.NoError(err)
}
