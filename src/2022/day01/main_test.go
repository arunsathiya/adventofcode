package main

import (
	"os"
	"testing"
)

func TestDay01Of2022(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test case from AOC",
			input:    "test.txt",
			expected: 24000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, err := os.Open(test.input)
			if err != nil {
				t.Fatalf("Failed to open file %v", err)
			}
			defer file.Close()
			output, err := Day01Of2022(file)
			if err != nil {
				t.Errorf("Expected %d, got %v", test.expected, err)
			}
			if output != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}
