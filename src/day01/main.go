package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func numbersFromWords(line string) string {
	numberRegex := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
	replacer := func(s string) int {
		replacements := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
			"zero":  0,
		}
		return replacements[s]
	}
	newText := numberRegex.ReplaceAllStringFunc(line, func(s string) string {
		return fmt.Sprint(replacer(s))
	})
	return newText
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibration := make([]int, 0)
		line = numbersFromWords(line)
		for _, v := range line {
			v, err := strconv.Atoi(string(v))
			if len(calibration) == 2 {
				break
			}
			if err == nil {
				calibration = append(calibration, v)
			}
		}
		switch len(calibration) {
		case 1:
			sum += calibration[0]*10 + calibration[0]
		case 2:
			sum += calibration[0]*10 + calibration[1]
		default:
			sum += 0
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
