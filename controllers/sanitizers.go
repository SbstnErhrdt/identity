package controllers

import "strings"

// SanitizeEmail sanitizes an email address
func SanitizeEmail(emailAddress string) string {
	return strings.TrimSpace(strings.ToLower(emailAddress))
}
