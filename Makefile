MAKEFLAGS += --silent

## generate: Generates the client code.
generate:
	openapi-generator-cli generate -i schema.yaml -g go -o ./k6 --git-user-id grafana --git-repo-id k6-cloud-openapi-client --package-name k6 -p isGoSubmodule=true -p disallowAdditionalPropertiesIfNotPresent=false -p withGoMod=false -t ./templates
	find k6 -name \*.go -exec goimports -w {} \;

## format: Applies Go formatting to code.
format:
	find k6 -name \*.go -exec goimports -w {} \;

## update-schema: Retrieves the latest version of the OpenAPI schema.
update-schema:
	wget -O schema.yaml https://api.k6.io/cloud/v6/openapi

## help: Prints a list of available build targets.
help:
	echo "Usage: make <OPTIONS> ... <TARGETS>"
	echo ""
	echo "Available targets are:"
	echo ''
	sed -n 's/^##//p' ${PWD}/Makefile | column -t -s ':' | sed -e 's/^/ /'
	echo
	echo "Targets run by default are: `sed -n 's/^all: //p' ./Makefile | sed -e 's/ /, /g' | sed -e 's/\(.*\), /\1, and /'`"

.PHONY: generate