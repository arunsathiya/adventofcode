package main

import (
	"os"
	"testing"
)

func TestDay03Of2022(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectedA int
		expectedB int
	}{
		{
			name:      "test case from AOC",
			input:     "test.txt",
			expectedA: 2,
			expectedB: 4,
		},
	}
	for _, test := range tests {
		input, err := os.Open(test.input)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		defer input.Close()
		pairs, err := parseInput(input)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		resultA, err := Day04Of2022PartA(pairs)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if resultA != test.expectedA {
			t.Errorf("expected %d, got %d", test.expectedA, resultA)
		}
		resultB, err := Day04Of2022PartB(pairs)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if resultB != test.expectedB {
			t.Errorf("expected %d, got %d", test.expectedB, resultB)
		}
	}
}
