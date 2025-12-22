# Download project dependencies
go mod tidy

# Install swag CLI if missing
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs
swag init -g cmd/server/main.go

# Set server port
1. Linux \ OSx:  
PORT=8080 go run cmd/server/main.go
2. Windows: \
$env:PORT="8080" \
go run cmd/server/main.go
3. Windows (CMD): \
set PORT=8080 \
go run cmd/server/main.g
4. IntelliJ: \
Run → Edit Configurations \ 
Select your Go run config \
Add: \
Environment variables: \
PORT=8080

# Run server
go run cmd/server/main.go

# Server address:
http://localhost:8080

# Swagger docs:
http://localhost:8080/swagger/index.html#