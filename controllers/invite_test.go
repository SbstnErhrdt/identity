package controllers

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInviteUser(t *testing.T) {
	ass := assert.New(t)

	conferencesUID := uuid.MustParse("721847d7-980a-4b1b-8b41-b1b53cab1400")
	riseUID := uuid.MustParse("a9e1767e-b3ec-455d-8701-348c3c1026f4")

	err := InviteUser(s, conferencesUID, &riseUID, "Erhardt", "Invitation", "Sebastian", "Erhardt", "test@erhardt.net", "https://erhardt.net")
	ass.NoError(err)

}
