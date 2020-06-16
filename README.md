# Account Gateway

Micro service to redirect requests

## Endpoints

### Activate account

`GET: /activate/{token}`

Http Status Code:

`HTTP/1.1 200 OK`

Payload Errors:

```json
{
  "error_account_gateway": {
    "http_status_code": 400,
    "http_body": "{\"errors\":\"Token is expired\"}"
  }
}
```

### Reset password

`POST: /reset-password`

Request:

```json
{
  "email": "test@bedu.org"
}
```
Http Status Code:

`HTTP/1.1 200 OK`

Payload Errors:

```json
{
  "error_account_gateway": {
    "http_status_code": 404,
    "http_body": "{\"errors\":\"cannot find user from email: \\\"test@bedu.org\\\"\"}"
  }
}
```

### Change Passsword

`POST: /change-password`

Request:

```json
{
  "email": "test@bedu.org",
  "password": "test"
}
```

Http Status Code:

`HTTP/1.1 200 OK`

Payload Errors:

```json
{
  "error_account_gateway": {
    "http_status_code": 400,
    "http_body": "{\"errors\":\"password: the length must be between 16 and 64.\"}"
  }
}
```
