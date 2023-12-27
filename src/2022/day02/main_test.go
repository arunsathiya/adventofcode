package main

import (
	"os"
	"testing"
)

func TestDay02Of2022(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectedA int
	}{
		{
			name:      "test case from AOC",
			input:     "test.txt",
			expectedA: 15,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, err := os.Open(test.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			defer file.Close()
			rounds, err := parseInput(file)
			if err != nil {
				t.Fatalf(err.Error())
			}
			sumOfRounds, err := Day02Of2022PartA(rounds)
			if err != nil {
				t.Errorf("Expected %d, got %v", test.expectedA, err)
			}
			if sumOfRounds != test.expectedA {
				t.Errorf("Expected %d, got %d", test.expectedA, sumOfRounds)
			}
		})
	}
}
