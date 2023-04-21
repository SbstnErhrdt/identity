package email

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDefaultRegistrationEmailResolver(t *testing.T) {
	ass := assert.New(t)
	email := "test@test.local"
	token := "tokenXYZ"
	res := DefaultRegistrationEmailResolver("test.local", email, token)
	html, err := res.Content()
	ass.NoError(err)
	ass.True(strings.Contains(html, email))
	ass.True(strings.Contains(html, token))
}
