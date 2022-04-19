# Identity

[![Go Report Card](https://goreportcard.com/badge/github.com/SbstnErhrdt/identity)](https://goreportcard.com/badge/github.com/SbstnErhrdt/identity)
[![Go Reference](https://pkg.go.dev/badge/github.com/SbstnErhrdt/identity.svg)](https://pkg.go.dev/github.com/SbstnErhrdt/identity)


An identity management system written in go using

* ORM (Object Relational Mapping) - Gorm
* JWT (JSON Web Token)
* Gin (HTTP framework)
* GraphQL

## Status

Under development

**⚠️ Experimental - Not ready for production.**

## Author

Sebastian Erhardt

## Environment Variables
```
SMPT_USER=user_name
SMPT_PASSWORD=secure_password
SMPT_SERVER=email-smtp.eu-central-1.amazonaws.com
SMPT_PORT=465
```

## Usage

```go
s := identity.NewService("APP", mail.Address{
    Name:    "App",
    Address: "no-reply@exameple.com",
}).
    SetSQLClient(connections.SQLClient).
    SetAuthConfirmationEndpoint("https://exameple.com/auth/confirm")
```

## Processes

### Registration

1. Check if identity already exists
2. Create new identity
3. Create confirmation token link
4. Send email / sms with confirmation link

```
Link endpoint 
Auth Confirmation Endpoint + /registration/{{Random Token}}
e.g. 
https://exameple.com/auth/confirm/registration/esrdzh534253qreafdsrgrqafeaar
```
5. Activate account

If the activation is expired, the account will be deleted.
The user will be able to register again.
### Invitation

1. Check if identity already exists
   1. If identity exists already:
      1. Create reference to entity
      2. send info email
   2. If identity does not exist:
      1. Create invitation token
      2. Create reference to entity
      3. Send email with invitation link
      4. Register with password

### Login

1. Check if identity exists
2. Generate token
3. Save ip and agent

### Lost Password

**OWASP Forgot Password Checklist**

The following short guidelines can be used as a quick reference to protect the forgot password service:

* Return a consistent message for both existent and non-existent accounts.
* Ensure that the time taken for the user response message is uniform.
* Use a side-channel to communicate the method to reset their password.
* Use URL tokens for the simplest and fastest implementation.
* Ensure that generated tokens or codes are:
  * Randomly generated using a cryptographically safe algorithm.
  * Sufficiently long to protect against brute-force attacks.
  * Stored securely.
  * Single use and expire after an appropriate period.
* Do not make a change to the account until a valid token is presented, such as locking out the account


### Process

1. 
2. Generate Token
3. Send Email with confirmation token link
4. Reset password

### Change Email

1. Generate Token
2. Send Email with confirmation token link to old email
3. Send Email with confirmation token link to new email
4. Change email
5. Send confirmation email


## Password

https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html

* Password Length
  * Minimum length of the passwords should be enforced by the application. Passwords shorter than 8 characters are considered to be weak (NIST SP800-63B).
  * Maximum password length should not be set too low, as it will prevent users from creating passphrases. A common maximum length is 64 characters due to limitations in certain hashing algorithms, as discussed in the Password Storage Cheat Sheet. It is important to set a maximum password length to prevent long password Denial of Service attacks.