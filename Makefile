SHELL := /bin/bash

# Define variables
AIR_BIN=$(shell go env GOPATH)/bin/air
BACKEND_DIR=backend
GOPATH_BIN=$(shell go env GOPATH)/bin

# Make sure the air binary exists
check-air-bin:
	@echo "Checking if air binary exists..."
	@if [ ! -f $(AIR_BIN) ]; then \
		echo "air binary not found. Installing air..."; \
		go install github.com/air-verse/air@latest; \
	else \
		echo "air binary already exists, skipping installation."; \
	fi

# Add the air alias to ~/.bashrc if it's not already there
add-air-alias:
	@echo "Checking if air alias is in bashrc..."
	@if ! grep -q "alias air='$(AIR_BIN)'" ~/.bashrc; then \
		echo "Adding air alias to bashrc..."; \
		echo "alias air='$(AIR_BIN)'" >> ~/.bashrc; \
	else \
		echo "air alias already exists in bashrc, skipping."; \
	fi

# Initialize air in the backend folder if .air.toml doesn't exist
init-air-backend:
	@echo "Checking if .air.toml exists in the backend..."
	@if [ ! -f $(BACKEND_DIR)/.air.toml ]; then \
		echo "Initializing air in the backend directory..."; \
		cd $(BACKEND_DIR) && $(AIR_BIN) init; \
	else \
		echo ".air.toml already exists in the backend, skipping initialization."; \
	fi

# Start docker-compose with --build in the go-server-reload directory
start-docker:
	@echo "Starting docker-compose with --build..."
	cd $(BACKEND_DIR) && docker compose -f docker-compose.go-reload.yml up --build

# Default target to run all steps sequentially
start-server-reload: check-air-bin add-air-alias init-air-backend start-docker
	@echo "Server reload process completed."

# Run Go application
run: start-server-reload

# Run Go mod tidy
tidy:
	go mod tidy

# Download Go mod dependencies
download:
	go mod download

# Serve Angular application (assuming this is required)
serve:
	ng serve

.PHONY: check-air-bin add-air-alias init-air-backend start-docker start-server-reload
