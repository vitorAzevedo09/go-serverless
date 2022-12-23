# go-serverless

url: https://u12kdqkwrh.execute-api.sa-east-1.amazonaws.com/staging

BUILD COMMAND:
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/main ./cmd/main.go
```

ZIP COMMAND:
```
zip -jrm ./build/main.zip ./build/main
```
