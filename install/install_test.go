package install

import (
	"github.com/SbstnErhrdt/env"
	"testing"
)

func TestInstall(t *testing.T) {
	env.LoadEnvFiles("../../../.env")
	Install()
}
