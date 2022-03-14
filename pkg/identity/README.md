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

1. Generate Token
2. Send Email with confirmation token link
3. Reset password

### Change Email

1. Generate Token
2. Send Email with confirmation token link to old email
3. Send Email with confirmation token link to new email
4. Change email
5. Send confirmation email