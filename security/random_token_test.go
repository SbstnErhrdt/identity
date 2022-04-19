package security

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	ass := assert.New(t)
	// Example: this will give us a 64 byte, base64 encoded output
	token, err := GenerateRandomString(64)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
	}
	ass.NoError(err)
	ass.NotEmpty(token)
	fmt.Println("token", token)
}
