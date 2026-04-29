MAKEFLAGS += --silent

## generate: Generates the client code.
generate:
	docker run --user 1000:1000 --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.9.0 generate -i /local/schema.yaml -g go -o /local/k6 --git-user-id grafana --git-repo-id k6-cloud-openapi-client --package-name k6 -p isGoSubmodule=true -p disallowAdditionalPropertiesIfNotPresent=false -p withGoMod=false -t /local/templates
	find k6 -name \*.go -exec goimports -w {} \;

## format: Applies Go formatting to code.
format:
	find k6 -name \*.go -exec goimports -w {} \;

## update-schema: Fetches the latest public API spec and merges it with provisioning.yaml.
update-schema:
	@if [ ! -f provisioning.yaml ]; then \
		echo "error: provisioning.yaml not found"; \
		exit 1; \
	fi
	wget -O public.yaml https://api.k6.io/cloud/v6/openapi
	yq eval-all -I 2 -c 'select(fileIndex == 0) *d select(fileIndex == 1)' provisioning.yaml public.yaml > schema.yaml
	rm -f public.yaml

## help: Prints a list of available build targets.
help:
	echo "Usage: make <OPTIONS> ... <TARGETS>"
	echo ""
	echo "Available targets are:"
	echo ''
	sed -n 's/^##//p' ${PWD}/Makefile | column -t -s ':' | sed -e 's/^/ /'
	@echo ""

.PHONY: generate
.DEFAULT_GOAL := help
