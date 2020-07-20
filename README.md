# Gateway

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
