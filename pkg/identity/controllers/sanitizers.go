package controllers

import "strings"

func SanitizeEmail(emailAddress string) string {
	return strings.TrimSpace(strings.ToLower(emailAddress))
}
