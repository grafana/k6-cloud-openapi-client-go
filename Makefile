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
	@echo "💡 Examples:"
	@echo "  make e2e-test                                      # Run all e2e tests"
	@echo "  make e2e-test TEST=TestAuthorizationAPI_Auth       # Run specific test"
	@echo "  make e2e-test ARGS=\"-timeout 30s -race\"            # Run with custom go test flags"
	@echo "  make e2e-test-coverage                             # Run e2e tests with function coverage"
	@echo "  make e2e-test-coverage MIN_COVERAGE=95             # Enforce minimum 95% function coverage"
	@echo "  make e2e-test-coverage TEST=TestSchedules          # Run specific test with coverage"
	@echo "  make e2e-test-coverage ARGS=\"-race\" MIN_COVERAGE=95  # With race detector"
	@echo ""

define TEST_SUMMARY_AWK
BEGIN { passed=0; failed=0; skipped=0; } \
{ print; } \
/^--- PASS:/ { passed++; pass[passed]=$$0; } \
/^--- FAIL:/ { failed++; fail[failed]=$$0; } \
/^--- SKIP:/ { skipped++; } \
END { \
	print ""; \
	print "📊 Test Results Summary:"; \
	print "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"; \
	print "  ✅ Passed:  " passed; \
	print "  ❌ Failed:  " failed; \
	print "  ⏭️  Skipped: " skipped; \
	if (failed > 0) { \
		print ""; \
		print "❌ Failed tests:"; \
		for (i=1; i<=failed; i++) { \
			gsub(/^--- FAIL: /, "  - ", fail[i]); \
			print fail[i]; \
		} \
	} \
	print "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"; \
	print ""; \
	if (failed > 0) exit 1; \
}
endef

e2e-test: ## Run tests (all tests or specific test with TEST=TestName, custom args with ARGS=...)
	@if [ -z "$(TEST)" ]; then \
		echo "🧪 Running all E2E tests..."; \
		go test -v $(ARGS) ./e2e/... 2>&1 | awk '$(TEST_SUMMARY_AWK)'; \
	else \
		echo "🧪 Running test: $(TEST)"; \
		go test -v -run "^$(TEST)$$" $(ARGS) ./e2e/... 2>&1 | awk '$(TEST_SUMMARY_AWK)'; \
	fi

e2e-test-coverage: ## Run e2e tests with function coverage (set MIN_COVERAGE=N to enforce minimum %, TEST/ARGS same as e2e-test)
	@$(MAKE) e2e-test ARGS="-coverpkg github.com/grafana/k6-cloud-openapi-client-go/k6 -coverprofile coverage.out $(ARGS)" TEST="$(TEST)"
	@echo ""
	@echo "📊 Analyzing coverage..."
	@echo ""
	@TOTAL=$$(go tool cover -func coverage.out | awk '$$1 ~ /k6\/api_/ && $$2 !~ /^(String|MarshalJSON|UnmarshalJSON|Error|GoString|init)$$/' | wc -l | tr -d ' '); \
	COVERED=$$(go tool cover -func coverage.out | awk '$$1 ~ /k6\/api_/ && $$2 !~ /^(String|MarshalJSON|UnmarshalJSON|Error|GoString|init)$$/ && $$3 != "0.0%"' | wc -l | tr -d ' '); \
	if [ "$$TOTAL" -eq 0 ]; then \
		echo "❌ Error: No API functions found in coverage report"; \
		exit 1; \
	fi; \
	FUNC_COV=$$(echo "scale=1; $$COVERED * 100 / $$TOTAL" | bc); \
	echo "📈 Coverage Summary:"; \
	echo "  API Function Coverage: $${FUNC_COV}% ($$COVERED/$$TOTAL functions)"; \
	echo ""; \
	UNCOVERED=$$(echo "$$TOTAL - $$COVERED" | bc); \
	if [ "$$UNCOVERED" -gt 0 ]; then \
		echo "🔍 Missing coverage ($$UNCOVERED functions):"; \
		go tool cover -func coverage.out | awk '$$1 ~ /k6\/api_/ && $$3 == "0.0%" && $$2 !~ /^(String|MarshalJSON|UnmarshalJSON|Error|GoString|init)$$/ {print "  - " $$1 ":" $$2}'; \
		echo ""; \
	fi; \
	if [ -n "$(MIN_COVERAGE)" ]; then \
		MEETS_MIN=$$(echo "$$FUNC_COV >= $(MIN_COVERAGE)" | bc -l); \
		if [ "$$MEETS_MIN" -eq 0 ]; then \
			echo "❌ Function coverage $${FUNC_COV}% is below minimum $(MIN_COVERAGE)%"; \
			exit 1; \
		else \
			echo "✅ Function coverage $${FUNC_COV}% meets minimum $(MIN_COVERAGE)%"; \
		fi; \
	fi

.PHONY: generate e2e-test-coverage
.DEFAULT_GOAL := help
