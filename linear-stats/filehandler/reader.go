package filehandler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ANSI color codes for warnings
const (
	colorBrightGold = "\033[93m"
	colorReset      = "\033[0m"
)

// ReadNumbersFromFile reads a file and extracts numeric values
func ReadNumbersFromFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty and commented lines
		if line == "" || line[0] == '#' {
			continue
		}

		// Remove quotes and commas
		cleanedLine := strings.ReplaceAll(line, `"`, "")
		cleanedLine = strings.ReplaceAll(cleanedLine, "'", "")
		cleanedLine = strings.ReplaceAll(cleanedLine, ",", "")

		// Parse number
		num, err := strconv.ParseFloat(cleanedLine, 64)
		if err != nil {
			warningMsg := fmt.Sprintf("Warning: Skipping invalid number on line %d: %s", lineNumber, line)
			underlineDots := strings.Repeat(".", len(warningMsg))
			fmt.Printf("%s%s%s\n%s%s%s\n", colorBrightGold, warningMsg, colorReset, colorBrightGold, underlineDots, colorReset)
			continue
		}

		numbers = append(numbers, num)
	}

	// Check for file reading errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return numbers, nil
}
