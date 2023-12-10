package main

import (
	"os"
	"testing"
)

func TestNumber(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"eightawjhkawaw", 88},
		{"8aw", 88},
	}
	for _, tt := range tests {
		got, err := number(tt.input)
		if err != nil {
			t.Errorf("%q returned an error %v", tt.input, err)
		}
		if got != tt.output {
			t.Errorf("Wanted: %d, Got: %d", tt.output, got)
		}
	}
}

func TestCalculateSum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		err      bool
	}{
		{"valid input", "a\n12", 12, false},
		{"valid input", "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", 281, false},
		{"valid input", "twothree\nfour", 67, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.WriteString(test.input); err != nil {
				tmpfile.Close()
				t.Fatal(err)
			}

			file, err := os.Open(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			result, err := calculateSum(file)
			if (err != nil) != test.err {
				t.Errorf("Test %s failed, expected error %v, got %v", test.name, test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
