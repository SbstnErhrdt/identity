package controllers

// InitializePasswordRecovery inits the password recovery process
func InitializePasswordRecovery(emailAddress string) (err error) {
	// sanitize email address
	emailAddress = SanitizeEmail(emailAddress)
	// create a new password recovery token
	// todo
	// create a entry in the database
	// todo
	// send email
	// todo
	return
}

func ChangePassword(token string, newPassword string) (err error) {
	// check if the token is valid
	// todo
	// check if the entry in the database is still valid

	// change the password
	// todo
	// mark the token as used
	// todo
	return
}
