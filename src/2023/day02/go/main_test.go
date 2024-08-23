package main

import (
	"os"
	"reflect"
	"testing"
)

func TestIdentifyCubes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output map[int]bool
		err    bool
	}{
		{"valid input", "Game 1: 7 blue, 6 green, 3 red; 3 red, 5 green, 1 blue; 1 red, 5 green, 8 blue; 3 red, 1 green, 5 blue\nGame 2: 9 green, 1 blue, 12 red; 1 blue, 18 green, 8 red; 2 blue, 6 green, 13 red; 3 blue, 13 red, 7 green; 5 blue, 4 red, 4 green; 6 blue, 7 green, 4 red\nGame 3: 5 blue, 9 red, 14 green; 10 green, 3 blue; 11 red, 2 blue, 8 green; 5 red, 2 blue; 5 blue, 7 green, 8 red\nGame 4: 2 red, 3 blue, 2 green; 17 green, 6 blue, 1 red; 3 blue, 5 green, 1 red; 4 red, 1 blue, 16 green; 5 red, 4 blue, 13 green; 14 green, 5 blue, 6 red\nGame 5: 3 red, 17 green, 10 blue; 9 blue, 5 green; 14 green, 9 blue, 11 red\nGame 6: 4 green, 18 blue, 3 red; 6 green, 8 blue, 9 red; 4 green, 9 blue, 7 red; 9 red, 1 green, 12 blue\nGame 7: 1 blue, 14 green; 1 red, 4 blue, 15 green; 3 blue, 6 green; 3 blue, 2 green, 1 red; 1 red, 3 green, 1 blue\nGame 8: 10 red, 3 blue, 3 green; 5 blue, 7 red, 3 green; 3 red, 3 green, 11 blue; 1 red, 7 green, 10 blue; 13 blue, 5 green, 5 red; 1 green, 17 blue, 3 red\nGame 9: 1 blue, 6 green; 7 green, 2 red; 3 red, 2 green; 1 blue, 4 red, 3 green; 7 green, 1 blue, 1 red\nGame 10: 14 green, 6 blue, 1 red; 8 green, 5 red, 1 blue; 8 green, 5 blue, 5 red; 2 green, 3 blue, 5 red", map[int]bool{
			1:  true,
			2:  false,
			3:  false,
			4:  false,
			5:  false,
			6:  false,
			7:  false,
			8:  false,
			9:  true,
			10: false,
		}, false},
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

			result, err := identifyCubes(file)
			if (err != nil) != test.err {
				t.Errorf("Test %s failed, expected error %v, got %v", test.name, test.err, err)
			}
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected %v, got %v", test.output, result)
			}
		})
	}
}
