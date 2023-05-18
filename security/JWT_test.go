package security

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testUID = uuid.New()

func TestGenerateToken(t *testing.T) {
	// Test env
	ass := assert.New(t)
	// subject uid

	// payload
	payload := map[string]interface{}{
		"test": "test",
	}
	// Create token
	tokenString, tokenUUUID, err := GenerateJWTToken(testUID, "test_audience", payload)
	// check
	ass.NoError(err)
	ass.NotEmpty(tokenString)
	ass.NotEmpty(tokenUUUID)
	fmt.Println(tokenString)
}

func TestParseToken(t *testing.T) {
	// Test env
	ass := assert.New(t)
	// Valid Token
	payload := map[string]interface{}{
		"test": "test",
	}
	// Create token
	tokenString, tokenUUUID, err := GenerateJWTToken(testUID, "test_audience", payload)
	ass.NoError(err)
	ass.NotEmpty(tokenUUUID)
	payload, err = ParseToken(tokenString)
	ass.NoError(err)
	// Invalid Token
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzUxOTQxNDcsImlhdCI6MT6zMzk4NDU0NywiaXNzdWVyIjoiR05PU0lTIiwicGF5bG9hZCI6eyJ1c2VybmFtZSI6Im1heF9tdXN0ZXJtYW5uIn0sInN1YmplY3QiOiJHTk9TSVMtQVBQIn0.__sn4ml2TZ7BuJAirirUZh6t2smcd2vA0dH4gaQ4EVU"
	payload, err = ParseToken(tokenString)
	ass.Error(err)
	ass.Empty(payload["subject"])
	ass.Empty(payload["issuer"])
	// Roles
	payload = map[string]interface{}{
		"admin": true,
	}
	// With attributes
	tokenString, tokenUUUID, err = GenerateJWTToken(testUID, "test_audience", payload)
	ass.NoError(err)
	ass.NotEmpty(tokenUUUID)
	res, err := ParseToken(tokenString)
	ass.NoError(err)
	ass.Equal(testUID.String(), res["sub"])
	ass.Equal("ERHARDT", res["iss"])
	p := res["payload"].(map[string]interface{})
	ass.True(p["admin"].(bool))
}
