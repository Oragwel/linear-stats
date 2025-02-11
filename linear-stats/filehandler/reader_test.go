package filehandler

import (
	"os"
	"testing"
)

// Helper function to create test files
func createTestFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func TestReadNumbersFromFile(t *testing.T) {
	// Define test files
	validFile := "../testdata/valid_data.txt"
	invalidFile := "../testdata/invalid_data.txt"

	// Create test files
	_ = createTestFile(validFile, "10\n20\n30\n40\n50\n")
	_ = createTestFile(invalidFile, "10\nhello\n30\n# comment\n50\n")

	// Test valid file
	numbers, err := ReadNumbersFromFile(validFile)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	expectedLen := 5
	if len(numbers) != expectedLen {
		t.Errorf("Expected %d numbers, got %d", expectedLen, len(numbers))
	}

	// Test invalid file (should skip invalid lines but still process valid numbers)
	numbers, err = ReadNumbersFromFile(invalidFile)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	expectedLen = 3 // "hello" is skipped, "# comment" is ignored
	if len(numbers) != expectedLen {
		t.Errorf("Expected %d numbers, got %d", expectedLen, len(numbers))
	}

	// Clean up test files
	_ = os.Remove(validFile)
	_ = os.Remove(invalidFile)
}
