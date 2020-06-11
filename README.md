# Account Gateway

Micro service to redirect requests

## Endpoints

### Activate account

`GET: /activate/{token}`

### Reset password

`POST: /reset-password`
```json
{
  "email": "test@bedu.org"
}
```

### Change Passsword

`POST: /change-password`
```json
{
  "email": "test@bedu.org",
  "password": "test"
}
```
