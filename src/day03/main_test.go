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
		{"valid input", "................................................965..583........389.................307.................512......................395.....387\n........................#....374...382....250...*..........737*....*896.395...........*....................$.........................#......\n..494.........532-...474......*.......#....*...................522......*..........%...........................%...+................269.....\n.....*..#................506..143........375......77.....155...........400.518...64....773...718..797........694....972.603.....*...........\n....479.795...............*..........800...........*.$.......264*636.......@..............&..*...*.......499...............*...5.20.........", 16337},
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
