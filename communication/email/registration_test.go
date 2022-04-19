package email

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestDefaultRegistrationEmailResolver(t *testing.T) {
	ass := assert.New(t)
	res := DefaultRegistrationEmailResolver("test.local", "test@test.local", "tokenXYZ")
	html, err := res.Content()
	ass.NoError(err)
	data := []byte(html)
	err = ioutil.WriteFile("./test_results/registration_confirmation.html", data, 0666)
	ass.NoError(err)
}
