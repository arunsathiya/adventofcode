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
			expectedA: 157,
			expectedB: 70,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, err := os.Open(test.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			defer file.Close()
			rucksacks, err := parseInput(file)
			if err != nil {
				t.Fatalf(err.Error())
			}
			priorities, err := Day03Of2022PartA(rucksacks)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if priorities != test.expectedA {
				t.Errorf("Expected %d, got %d", test.expectedA, priorities)
			}
			file.Seek(0, 0)
			prioritiesPartB, err := Day03Of2022PartB(rucksacks)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if prioritiesPartB != test.expectedB {
				t.Errorf("Expected %d, got %d", test.expectedB, prioritiesPartB)
			}
		})
	}
}
