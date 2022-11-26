# Templates

## Import
- helper
  - `go get github.com/joho/godotenv`
- iam
  - `go get github.com/zitadel/oidc/pkg/oidc`
  - `go get github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management`
  - `go get github.com/zitadel/zitadel-go/v2/pkg/client/management`
  - `go get github.com/zitadel/zitadel-go/v2/pkg/client/zitadel`
- logging
  - `go get github.com/getsentry/sentry-go`

## ENV

- helper
  - nothing
- iam
  - `ZITADEL_ISSUER`
  - `ZITADEL_API`
  - `ZITADEL_ISSUER`
  - `ZITADEL_ISSUER`
- logging
  - `APP_ENV`
  - `LOGGING_DNS`

## ToDo

- [ ] refactor logging-functions -> the most of the Error and Info-Function are identical