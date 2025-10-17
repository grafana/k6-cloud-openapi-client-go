# Tools

This directory contains development tools for the k6 Cloud OpenAPI Client.

## API Coverage Analyzer

The `api_coverage_analyzer.go` script analyzes which API functions from the `k6/` directory are actually called in the e2e tests.

### Usage

```bash
# Basic usage - just report coverage
make api-coverage

# Enforce minimum coverage threshold (fails if below)
make api-coverage MIN_COVERAGE=95

# Run directly with Go
go run tools/api_coverage_analyzer.go

# Run with minimum coverage requirement
go run tools/api_coverage_analyzer.go -min-coverage=90.0
```

**Exit Codes:**
- `0`: Success (coverage meets threshold if specified)
- `1`: Failure (coverage below threshold, validation error, or parsing error)

### What it Does

1. **Discovers API Functions**: Parses all `api_*.go`, `client.go`, and `configuration.go` files to find exported functions
2. **Finds Function Calls**: Analyzes e2e test files to find which functions are called
3. **Matches & Reports**: Shows which API functions are covered by tests and which are not

### Skip Patterns

The analyzer automatically skips common boilerplate functions that don't need explicit test coverage.

**Global Patterns** (apply to all files, key `"*"`):
- Getters (`GetId`, `GetName`, etc.)
- Setters (`SetId`, `SetName`, etc.)
- Has checks (`HasId`, etc.)
- Ok methods (`GetIdOk`, etc.)
- Constructors (`NewModel`, etc.)
- Execute methods (`AuthExecute`, etc.)
- JSON marshaling/unmarshaling
- ToMap converters

**File-Specific Patterns**:
- `client.go`: Utility methods like `CacheExpires`, `Model`
- `configuration.go`: Internal methods like `ServerURL`, `AddDefaultHeader`

These patterns are defined in `api_coverage_analyzer.go` and can be customized by editing the `skipPatterns` map:

```go
var skipPatterns = map[string][]string{
    "*": {              // Global patterns
        `^Get[A-Z].*`,  // Getters
        `^Set[A-Z].*`,  // Setters
        // ... more patterns
    },
    "client.go": {      // File-specific patterns
        `^CacheExpires$`,
        `^Model$`,
    },
    "your_file.go": {   // Add your own file-specific patterns
        `^YourFunction$`,
    },
}
```

### Output

The analyzer provides:
- **Summary**: Total functions, called percentage, uncalled count
- **Detailed Report**: Lists uncalled functions grouped by file with line numbers

### Example Output

**Without minimum coverage:**
```
📊 API Function Coverage Summary:
  Total API functions:     130
  Called in tests:         130 (100.0%)
  Not called:              0 (0.0%)
  Skipped (patterns):      94
  Minimum required:        Not set (reporting mode)

❌ Uncalled API Functions by File:
api_authorization.go:
  ✅ (all functions covered)
...
```

**With minimum coverage (95%):**
```
📊 API Function Coverage Summary:
  Total API functions:     130
  Called in tests:         130 (100.0%)
  Not called:              0 (0.0%)
  Skipped (patterns):      94
  Minimum required:        95.0%

✅ Coverage 100.0% meets minimum threshold of 95.0%
```

### How It Works

The analyzer uses Go's AST (Abstract Syntax Tree) parser to:

1. **Parse API files** and extract function declarations with their receiver types
2. **Parse test files** and extract function calls (including chained calls)
3. **Match calls to declarations** by function name and receiver type
4. **Apply skip patterns** using regex to filter out boilerplate
5. **Generate a report** showing coverage statistics

This provides a more meaningful coverage metric than traditional line coverage for OpenAPI-generated client code, focusing on which API operations are actually tested.

