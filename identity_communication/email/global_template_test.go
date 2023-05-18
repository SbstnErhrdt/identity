package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalTemplate_Style(t *testing.T) {
	ass := assert.New(t)
	res := DefaultGlobalTemplate.Style()
	ass.NotEmpty(res)
}
