.PHONY: help install test docs
# This Makefile is designed to automate the build process for a project.
# It includes targets for compiling the code, cleaning up build artifacts,
# and other common tasks. The default target is set to 'help', which provides
# a summary of available targets and their descriptions.
#
# Targets:
# - install: Installs the built executables and other necessary files to the system.
# - test: Runs tests for the project.
# - docs: Generates documentation for the project.

install: ## Install the plugin
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/chussenot/backstage.plugin *.go

test:
	go test -v ./...

docs:
	go generate ./...

config: ## Copy the backstage config to the steampipe config folder
	cp .steampipe/config/backstage.spc $(HOME)/.steampipe/config/backstage.spc
