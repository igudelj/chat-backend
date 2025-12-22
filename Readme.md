# Download project dependencies
go mod tidy

# Install swag CLI if missing
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs
swag init -g cmd/server/main.go

# Run server
go run cmd/server/main.go

# Server address:
http://localhost:8080

# Swagger docs:
http://localhost:8080/swagger/index.html#