BINARY_NAME=renco-boilerplate
build:
	@go build -o bin/${BINARY_NAME} main.go

run-http:
	@./bin/${BINARY_NAME} http

install:
	@echo "Installing dependencies...."
	@rm -rf vendor
	@rm -f Gopkg.lock
	@rm -f glide.lock
	@go mod tidy && go mod download && go mod vendor

start-http:
	@go run main.go http
	
# Define the base command for migration
MIGRATE_CMD=go run main.go db:migrate

# Default target
.DEFAULT_GOAL := help

# Show help message
help:
	@echo "Usage:"
	@echo "  make migrate-status   # Show migration status"
	@echo "  make migrate-up       # Apply all pending migrations"
	@echo "  make migrate-redo     # Redo the most recent migration"
	@echo "  make migrate-reset    # Rollback all migrations and reapply"
	@echo "  make migrate-fix      # Fix broken migrations"
	@echo "  make migrate-create   # Create a new migration with a specified name"
	@echo "  make migrate-create NAME=<migration_name> [TYPE=sql] # Create a new migration with specified name and type"

# Show migration status
migrate-status:
	@$(MIGRATE_CMD) status

# Apply all pending migrations
migrate-up:
	@$(MIGRATE_CMD) up

# Redo the most recent migration
migrate-redo:
	@$(MIGRATE_CMD) redo

# Rollback all migrations and reapply
migrate-reset:
	@$(MIGRATE_CMD) reset

# Fix broken migrations
migrate-fix:
	@$(MIGRATE_CMD) fix

# Create a new migration
migrate-create:
	@$(MIGRATE_CMD) create ${NAME} ${TYPE}

# Usage examples:
# To create a migration named 'users' with SQL type:
# make migrate-create NAME=users TYPE=sql
#
# To create a migration named 'orders' with default type (if TYPE is not provided):
# make migrate-create NAME=orders
