package main

import (
	"fmt"
	"os"
	"path/filepath"

	"linear-stats/filehandler"
	"linear-stats/mathutils"
)

func main() {
	// Check for correct arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	// Get file path and validate extension
	filePath := os.Args[1]
	if filepath.Ext(filePath) != ".txt" {
		fmt.Println("Error: File must have a .txt extension")
		return
	}

	// Read numbers from the file
	yValues, err := filehandler.ReadNumbersFromFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Ensure valid data is present
	if len(yValues) == 0 {
		fmt.Println("No valid data found in file")
		return
	}

	// Compute regression and correlation
	slope, intercept, pearson := mathutils.ComputeRegression(yValues)

	// Print results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
}
