# Copyright (c) 2020 Contaim, LLC
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

# Constants
PROJECTNAME = response
PROJECT_IMPORT = "github.com/responserms/server"
CORE_PACKAGE_PATH = "response"
CORE_LD_PATH = "$(PROJECT_IMPORT)/$(CORE_PACKAGE_PATH)"

# Build-time variables that are injected
LDFLAGS=-ldflags "-X '$(CORE_LD_PATH).gitTag=`git describe --tags --always`' -X '$(CORE_LD_PATH).gitCommit=$(BUILD)'"

# # load in the .env file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

bootstrap: # bootstrap the build by downloading additional tools that may be used by devs
	go generate -tags tools tools/tools.go

generate:
	@sh scripts/generate.sh

docs-dev:
	cd docs/ && yarn dev && cd ..

gen:
	@go generate ./...

gen-ent:
	@go run ./ent/cmd/generate/main.go

gen-graphql:
	@gqlgen generate

lint:
	golangci-lint run -v ./...