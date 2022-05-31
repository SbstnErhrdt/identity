package identity_controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	ass := assert.New(t)

	err := Register(
		s,
		"test@erhardt.net",
		"MaggieThatcherIs110%Sexy!",
		true,
		"test-agent",
		"0.0.0.0",
		"dev.local",
	)
	if err != nil {
		ass.NoError(err)
		return
	}
}

func TestLogin(t *testing.T) {
	ass := assert.New(t)
	token, err := Login(s, "", "", "", "")
	ass.NoError(err)
	ass.NotEmpty(token)
}
