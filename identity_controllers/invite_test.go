package identity_controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInviteUser(t *testing.T) {
	ass := assert.New(t)
	err := InviteUser(s, "rise-workshop.com", "Invitation", "Sebastian", "Erhardt", "test@erhardt.net", "https://erhardt.net")
	ass.NoError(err)
}