# Identity

An identity management system written in go using

* ORM (Object Relational Mapping) - Gorm
* JWT (JSON Web Token)
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