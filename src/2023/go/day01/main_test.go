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
		{"two1nine", 11},
		{"eightwothree", 0},
		{"abcone2threexyz", 22},
		{"xtwone3four", 33},
		{"4nineeightseven2", 42},
		{"zoneight234", 24},
		{"7pqrstsixteen", 77},
		{"eightawjhkawaw", 0},
		{"8aw", 88},
		{"76xkqjzqtwonfour", 76},
		{"sixthree8sixjxjqsjgjgp", 88},
		{"38bgcczgtninefivefive", 38},
		{"sixthree4eight", 44},
		{"nhp3zdc", 33},
		{"279four", 29},
		{"vzxf4tqrljgxmthreejcr", 44},
		{"bbm4twoeight8oneone3one", 43},
		{"nineninesix6nine", 66},
		{"fourseven5seveneightsvtkcjdrfour", 55},
		{"eightwo", 0},
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
		{"valid input", "a\n12\neightwo", 12, false},
		{"valid input", "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", 209, false},
		{"valid input", "twothree\nfour", 0, false},
		{"valid input", "76xkqjzqtwonfour\nsixthree8sixjxjqsjgjgp\n38bgcczgtninefivefive\nsixthree4eight\nnhp3zdc\n279four\nvzxf4tqrljgxmthreejcr\nbbm4twoeight8oneone3one\nnineninesix6nine\nfourseven5seveneightsvtkcjdrfour", 516, false},
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
