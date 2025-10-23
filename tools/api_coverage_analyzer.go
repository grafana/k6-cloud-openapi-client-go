package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Skip patterns for generated boilerplate functions
// Key "*" applies to all files (global patterns)
// Specific filenames apply only to that file
var skipPatterns = map[string][]string{
	"*": { // Global patterns (apply to all files)
		`^Get[A-Z].*`,      // Getters like GetId, GetName
		`^Set[A-Z].*`,      // Setters like SetId, SetName
		`^Has[A-Z].*`,      // Has checks like HasId
		`.*Ok$`,            // GetXXXOk methods
		`^New[A-Z].*`,      // Constructors like NewModel
		`.*Execute$`,       // Low-level execute methods
		`^Marshal.*`,       // JSON marshaling
		`^Unmarshal.*`,     // JSON unmarshaling
		`^.*ToMap$`,        // ToMap converters
		`^Is[A-Z].*`,       // IsSet, IsNil, etc.
		`^Unset$`,          // Unset methods
		`Ptr$`,             // Pointer helpers
		`^.*WithDefaults$`, // WithDefaults constructors
	},
	"client.go": {
		`^CacheExpires$`, // Skip CacheExpires utility
		`^Model$`,        // Skip Model utility
	},
	"configuration.go": {
		`^String$`,               // Skip String utility
		`^AddDefaultHeader$`,     // Skip AddDefaultHeader (rarely used)
		`^URL$`,                  // Skip URL utility
		`^ServerURL$`,            // Skip ServerURL (internal)
		`^ServerURLWithContext$`, // Skip ServerURLWithContext (internal)
	},
}

var skipRegexes map[string][]*regexp.Regexp

// Command line flags
var (
	minCoverage = flag.Float64("min-coverage", 0.0, "Minimum coverage percentage required (0-100). Fails if below this threshold.")
)

// APIFunction represents a function declaration in API files
type APIFunction struct {
	Name       string
	File       string
	Line       int
	RecvType   string // For methods: the receiver type
	IsCalled   bool
	ShouldSkip bool
}

func init() {
	// Compile skip patterns for each file (including global "*")
	skipRegexes = make(map[string][]*regexp.Regexp)

	for file, patterns := range skipPatterns {
		for _, pattern := range patterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: invalid regex pattern %q for file %q: %v\n", pattern, file, err)
				continue
			}
			skipRegexes[file] = append(skipRegexes[file], re)
		}
	}
}

func main() {
	flag.Parse()

	// Validate min-coverage flag
	if *minCoverage < 0 || *minCoverage > 100 {
		fmt.Fprintf(os.Stderr, "Error: min-coverage must be between 0 and 100\n")
		os.Exit(1)
	}

	exitCode, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(exitCode)
}

func run() (int, error) {
	// Find project root
	projectRoot, err := findProjectRoot()
	if err != nil {
		return 1, fmt.Errorf("failed to find project root: %w", err)
	}

	k6Dir := filepath.Join(projectRoot, "k6")
	e2eDir := filepath.Join(projectRoot, "e2e")

	// Step 1: Discover all API functions
	apiFunctions, err := discoverAPIFunctions(k6Dir)
	if err != nil {
		return 1, fmt.Errorf("failed to discover API functions: %w", err)
	}

	// Step 2: Find all function calls in e2e tests
	calledFunctions, err := findCalledFunctions(e2eDir)
	if err != nil {
		return 1, fmt.Errorf("failed to find called functions: %w", err)
	}

	// Step 3: Mark functions as called
	markCalledFunctions(apiFunctions, calledFunctions)

	// Step 4: Generate report and check coverage threshold
	exitCode := generateReport(apiFunctions)

	return exitCode, nil
}

func findProjectRoot() (string, error) {
	// Start from current directory and look for go.mod
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find go.mod")
		}
		dir = parent
	}
}

func discoverAPIFunctions(k6Dir string) (map[string]*APIFunction, error) {
	functions := make(map[string]*APIFunction)
	fset := token.NewFileSet()

	// Files to scan
	targetFiles := []string{
		"api_authorization.go",
		"api_load_tests.go",
		"api_load_zones.go",
		"api_projects.go",
		"api_schedules.go",
		"api_test_runs.go",
		"client.go",
		"configuration.go",
	}

	for _, filename := range targetFiles {
		filePath := filepath.Join(k6Dir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			continue // Skip if file doesn't exist
		}

		file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", filename, err)
		}

		// Walk the AST to find function declarations
		ast.Inspect(file, func(n ast.Node) bool {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			// Only consider exported functions
			if !ast.IsExported(funcDecl.Name.Name) {
				return true
			}

			funcName := funcDecl.Name.Name
			recvType := ""

			// Get receiver type for methods
			if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
				recvType = exprToString(funcDecl.Recv.List[0].Type)
			}

			// Create unique key
			key := funcName
			if recvType != "" {
				key = recvType + "." + funcName
			}

			// Check if should skip (pass filename for file-specific patterns)
			shouldSkip := matchesSkipPattern(funcName, filename)

			functions[key] = &APIFunction{
				Name:       funcName,
				File:       filename,
				Line:       fset.Position(funcDecl.Pos()).Line,
				RecvType:   recvType,
				IsCalled:   false,
				ShouldSkip: shouldSkip,
			}

			return true
		})
	}

	return functions, nil
}

func findCalledFunctions(e2eDir string) (map[string]bool, error) {
	called := make(map[string]bool)
	fset := token.NewFileSet()

	err := filepath.Walk(e2eDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		// Walk AST to find function calls
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Extract function name from call expression
			funcName := extractFunctionName(callExpr.Fun)
			if funcName != "" {
				called[funcName] = true
			}

			return true
		})

		return nil
	})

	return called, err
}

func extractFunctionName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		// Direct function call: FunctionName()
		return e.Name

	case *ast.SelectorExpr:
		// Method call: obj.Method()
		// We want just the method name
		methodName := e.Sel.Name

		// Also try to get the full chain for matching
		// e.g., testClient.LoadTestsAPI.LoadTestsList
		var parts []string
		curr := expr
		for {
			switch x := curr.(type) {
			case *ast.SelectorExpr:
				parts = append([]string{x.Sel.Name}, parts...)
				curr = x.X
			case *ast.Ident:
				parts = append([]string{x.Name}, parts...)
				curr = nil
			default:
				curr = nil
			}
			if curr == nil {
				break
			}
		}

		// Return both the simple method name and the full chain
		// The matching logic will try both
		return methodName

	case *ast.CallExpr:
		// Chained call: obj.Method().AnotherMethod()
		return extractFunctionName(e.Fun)
	}

	return ""
}

func markCalledFunctions(apiFunctions map[string]*APIFunction, calledFunctions map[string]bool) {
	for key, apiFunc := range apiFunctions {
		// Check if function name is called (simple match)
		if calledFunctions[apiFunc.Name] {
			apiFunc.IsCalled = true
			continue
		}

		// Check if the full key matches (RecvType.FuncName)
		if calledFunctions[key] {
			apiFunc.IsCalled = true
			continue
		}

		// Check variations of receiver type matching
		if apiFunc.RecvType != "" {
			// Try without pointer: *LoadTestsAPIService -> LoadTestsAPIService
			cleanRecv := strings.TrimPrefix(apiFunc.RecvType, "*")
			if calledFunctions[cleanRecv+"."+apiFunc.Name] {
				apiFunc.IsCalled = true
				continue
			}
		}
	}
}

func matchesSkipPattern(funcName string, fileName string) bool {
	// Check global patterns (key "*")
	if globalPatterns, exists := skipRegexes["*"]; exists {
		for _, re := range globalPatterns {
			if re.MatchString(funcName) {
				return true
			}
		}
	}

	// Check file-specific patterns
	if filePatterns, exists := skipRegexes[fileName]; exists {
		for _, re := range filePatterns {
			if re.MatchString(funcName) {
				return true
			}
		}
	}

	return false
}

func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return "*" + exprToString(e.X)
	case *ast.SelectorExpr:
		return exprToString(e.X) + "." + e.Sel.Name
	default:
		return ""
	}
}

func generateReport(apiFunctions map[string]*APIFunction) int {
	// Group functions by file
	byFile := make(map[string][]*APIFunction)
	var totalFunctions, calledFunctions, skippedFunctions int

	for _, fn := range apiFunctions {
		byFile[fn.File] = append(byFile[fn.File], fn)
		totalFunctions++

		if fn.IsCalled {
			calledFunctions++
		}
		if fn.ShouldSkip {
			skippedFunctions++
		}
	}

	// Calculate uncalled (excluding skipped)
	uncalledCount := 0
	reportableFunctions := 0
	for _, fn := range apiFunctions {
		if !fn.ShouldSkip {
			reportableFunctions++
			if !fn.IsCalled {
				uncalledCount++
			}
		}
	}

	// Calculate coverage percentage
	coveragePercent := 0.0
	if reportableFunctions > 0 {
		coveragePercent = float64(reportableFunctions-uncalledCount) / float64(reportableFunctions) * 100
	}

	// Print summary
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📊 API Function Coverage Summary:")
	fmt.Printf("  Total API functions:     %d\n", reportableFunctions)
	fmt.Printf("  Called in tests:         %d (%.1f%%)\n", reportableFunctions-uncalledCount, coveragePercent)
	fmt.Printf("  Not called:              %d (%.1f%%)\n", uncalledCount, float64(uncalledCount)/float64(reportableFunctions)*100)
	fmt.Printf("  Skipped (patterns):      %d\n", skippedFunctions)

	// Show minimum coverage configuration
	if *minCoverage > 0 {
		fmt.Printf("  Minimum required:        %.1f%%\n", *minCoverage)
	} else {
		fmt.Println("  Minimum required:        Not set (reporting mode)")
	}

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("❌ Uncalled API Functions by File:")
	fmt.Println()

	// Sort files for consistent output
	var files []string
	for file := range byFile {
		files = append(files, file)
	}
	sort.Strings(files)

	// Print uncalled functions by file
	for _, file := range files {
		functions := byFile[file]

		// Sort functions by line number
		sort.Slice(functions, func(i, j int) bool {
			return functions[i].Line < functions[j].Line
		})

		// Collect uncalled functions (excluding skipped)
		var uncalled []*APIFunction
		for _, fn := range functions {
			if !fn.IsCalled && !fn.ShouldSkip {
				uncalled = append(uncalled, fn)
			}
		}

		// Print file section
		fmt.Printf("%s:\n", file)
		if len(uncalled) == 0 {
			fmt.Println("  ✅ (all functions covered)")
		} else {
			for _, fn := range uncalled {
				fmt.Printf("  - %s (line %d)\n", fn.Name, fn.Line)
			}
		}
		fmt.Println()
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Check if coverage meets minimum threshold
	if *minCoverage > 0 && coveragePercent < *minCoverage {
		fmt.Println()
		fmt.Printf("❌ Coverage %.1f%% is below minimum threshold of %.1f%%\n", coveragePercent, *minCoverage)
		fmt.Println()
		return 1 // Fail
	}

	if *minCoverage > 0 {
		fmt.Println()
		fmt.Printf("✅ Coverage %.1f%% meets minimum threshold of %.1f%%\n", coveragePercent, *minCoverage)
		fmt.Println()
	}

	return 0 // Success
}
