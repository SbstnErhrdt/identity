package identity_controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	ass := assert.New(t)
	in := "hello world"
	out := AnonymizeString(in)
	ass.NotEmpty(out)
	ass.NotEqual(in, out)
}
