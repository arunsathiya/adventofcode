package main

import (
	"os"
	"testing"
)

func TestDay01Of2022(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectedA int
		expectedB int
	}{
		{
			name:      "test case from AOC",
			input:     "test.txt",
			expectedA: 24000,
			expectedB: 45000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, err := os.Open(test.input)
			if err != nil {
				t.Fatalf("Failed to open file %v", err)
			}
			defer file.Close()
			blocks, err := parseInput(file)
			if err != nil {
				t.Fatalf(err.Error())
			}
			highestSum, err := Day01Of2022PartA(blocks)
			if err != nil {
				t.Errorf("Expected %d, got %v", test.expectedA, err)
			}
			if highestSum != test.expectedA {
				t.Errorf("Expected %d, got %d", test.expectedA, highestSum)
			}
			file.Seek(0, 0)
			sumOfThree, err := Day01Of2022PartB(blocks)
			if err != nil {
				t.Errorf("Expected %d, got %v", test.expectedB, err)
			}
			if sumOfThree != test.expectedB {
				t.Errorf("Expected %d, got %d", test.expectedB, highestSum)
			}
		})
	}
}
