package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
)

func numbersFromWords(line string) string {
	numberRegex := regexp2.MustCompile(`(?=(one|two|three|four|five|six|seven|eight|nine|zero))`, regexp2.None)
	newText, err := numberRegex.ReplaceFunc(line, func(m regexp2.Match) string {
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
		return fmt.Sprint(replacements[m.String()])
	}, -1, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newText)
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
