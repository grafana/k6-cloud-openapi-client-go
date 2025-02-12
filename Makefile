MAKEFLAGS += --silent

## generate: Generates the client code.
generate:
	docker run --user 1000:1000 --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.9.0 generate -i /local/schema.yaml -g go -o /local/k6 --git-user-id grafana --git-repo-id k6-cloud-openapi-client --package-name k6 -p isGoSubmodule=true -p disallowAdditionalPropertiesIfNotPresent=false -p withGoMod=false -t /local/templates
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