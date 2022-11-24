package identity_controllers

import "testing"

func TestHash(t *testing.T) {
	h := Hash("hello world")
	t.Log("hash", h)
}
