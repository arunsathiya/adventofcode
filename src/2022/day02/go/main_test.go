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
		expectedB int
	}{
		{
			name:      "test case from AOC",
			input:     "test.txt",
			expectedA: 15,
			expectedB: 12,
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
				t.Fatalf(err.Error())
			}
			if sumOfRounds != test.expectedA {
				t.Errorf("Expected %d, got %d", test.expectedA, sumOfRounds)
			}
			file.Seek(0, 0)
			partBSumOfRounds, err := Day02Of2022PartB(rounds)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if partBSumOfRounds != test.expectedB {
				t.Errorf("Expected %d, got %d", test.expectedB, partBSumOfRounds)
			}
		})
	}
}
