# Download project dependencies
go mod tidy

# Install swag CLI if missing
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs
swag init -g cmd/server/main.go

# Set environment
1. Linux \ OSx:  
APP_ENV=dev go run cmd/server/main.go
2. Windows: \
$env:APP_ENV="dev" \
go run cmd/server/main.go
3. Windows (CMD): \
set APP_ENV=dev \
go run cmd/server/main.g
4. IntelliJ: \
Run → Edit Configurations \ 
Select your Go run config \
Add: \
Environment variables: \
APP_ENV=dev

# Run server
go run cmd/server/main.go

# Server address:
http://localhost:8080

# Swagger docs:
http://localhost:8080/swagger/index.html#

## Parameters:
Port=xxxx (8080 used if none are passed or set in environment)
APP_ENV=xxx