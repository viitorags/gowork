.PHONY: docs clean

# Variables
APP_NAME=gowork

# Tasks
docs:
	@swag init -g main.go --dir ./ --output ./docs
clean: 
	@rm -rf $(APP_NAME)
	@rm -rf ./docs
