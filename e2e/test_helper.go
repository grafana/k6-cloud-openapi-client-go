package e2e

import (
	"fmt"
	"regexp"
)

// ExtractCoveragePercentage extracts the total coverage percentage from coverage output
// This is a utility function that can be used if needed for custom test runners
func ExtractCoveragePercentage(coverageOutput string) (float64, error) {
	re := regexp.MustCompile(`total:.*?(\d+\.\d+)%`)
	matches := re.FindStringSubmatch(coverageOutput)
	if len(matches) < 2 {
		return 0, fmt.Errorf("could not find total coverage in output")
	}

	var percentage float64
	_, err := fmt.Sscanf(matches[1], "%f", &percentage)
	if err != nil {
		return 0, fmt.Errorf("could not parse coverage percentage: %w", err)
	}

	return percentage, nil
}
