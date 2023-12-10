package main

import "testing"

func TestNumericalValues(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"two1nine", 21},
		{"eightwothree", 82},
		{"abcone2threexyz", 12},
		{"xtwone3four", 21},
		{"4nineeightseven2", 49},
		{"zoneight234", 18},
		{"7pqrstsixteen", 76},
	}
	for _, tt := range tests {
		got, err := numericalValues(tt.input)
		if err != nil {
			t.Errorf("%q returned an error %v", tt.input, err)
		}
		if got != tt.output {
			t.Errorf("Wanted: %d, Got: %d", tt.output, got)
		}
	}
}
