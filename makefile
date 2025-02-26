.PHONY: all build up down check-db run test swagger

# Application and database container names
APP_NAME=app
DB_CONTAINER_NAME=postgres

# Default target: build and start the application
all: build up

# Build the application
build:
	@echo "Building the Go application..."
	docker-compose build $(APP_NAME)

# Start the containers
up: check-db
	@echo "Starting the application..."
	docker-compose up -d $(APP_NAME)

# Stop the containers
down:
	@echo "Stopping all containers..."
	docker-compose down

# Check if the database container is already running
check-db:
	@echo "Checking if the database container is already running..."
	@if [ $$(docker ps -q -f name=$(DB_CONTAINER_NAME)) ]; then \
		echo "Database container is already running."; \
	else \
		echo "Starting the database container..."; \
		docker-compose up -d $(DB_CONTAINER_NAME); \
	fi

# Run the application
run:
	docker-compose -f docker-compose.yaml up -d
##
swagger:
swagger:
	export GOFLAGS="-mod=mod" && swag init --output ./docs --generalInfo ./cmd/main.go --parseInternal --parseDependency && export GOFLAGS="-mod=vendor"
test:
	@echo "Running tests..."
	go test ./integration_test/... -v -cover
