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

# liquibase database setup (optional):
1. download odbc driver: https://www.postgresql.org/download/
2. Change the location of the driver jar file appropriately. Current one was used due to the versions of the database etc. 
(line 28 in the docker-compose.yml)

# Start docker:
docker compose up -d

# Run server
go run cmd/server/main.go

# Server address:
http://localhost:8080

# Swagger docs address:
http://localhost:8080/swagger/index.html#

# Keycloak address:
http://localhost:8081
http://localhost:8081/admin

## Parameters:
Port=xxxx (8080 used if none are passed or set in environment)
APP_ENV=xxx