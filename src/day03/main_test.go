package main

import (
	"os"
	"testing"
)

func TestNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"valid input", "................................................96\n........................#....374...382....250...*.\n..494.........532-...474......*.......#....*......\n.....*..#................506..143........375......\n....479.795...............*..........800..........\n515...................512.484...*....*...=......39\n..........607...........&.....105...906...910.....\n..........*...781................................9\n.......850.......=........559..249................\n....................157....*...&....978...........", 11663},
		{"valid input", "........\n.24..4..\n......*.", 4},
		{"valid input", "12.......*..\n+.........34\n.......-12..\n..78........\n..*....60...\n78..........\n.......23...\n....90*12...\n............\n2.2......12.\n.*.........*\n1.1.......56", 413},
		{"valid input", "12.......*..\n+.........34\n.......-12..\n..78........\n..*....60...\n78.........9\n.5.....23..$\n8...90*12...\n............\n2.2......12.\n.*.........*\n1.1..503+.56", 925},
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

			sum, err := numbers(file)
			if err != nil {
				t.Errorf("Test %s failed with %v", test.name, err)
			}
			if sum != test.output {
				t.Errorf("Expected %d, got %d", test.output, sum)
			}
		})
	}
}
