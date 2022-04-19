package email

import (
	"fmt"
	"testing"
)

func TestGlobalTemplate_Style(t *testing.T) {
	res := DefaultGlobalTemplate.Style()
	fmt.Println(res)
}
