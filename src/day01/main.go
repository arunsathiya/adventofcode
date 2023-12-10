package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
)

func number(line string) (int, error) {
	re := regexp2.MustCompile(`(?=(one|two|three|four|five|six|seven|eight|nine|zero|1|2|3|4|5|6|7|8|9|0))`, regexp2.None)
	replacements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
	}
	firstTwoNumbers := ""
	match, err := re.FindStringMatch(line)
	if err != nil {
		return 0, err
	}
	for match != nil {
		number, ok := replacements[match.GroupByNumber(1).String()]
		if ok {
			firstTwoNumbers += number
		}
		if len(firstTwoNumbers) == 2 {
			break
		}
		match, err = re.FindNextMatch(match)
		if err != nil {
			return 0, err
		}
	}
	number := 0
	switch len(firstTwoNumbers) {
	case 0:
		return 0, nil
	case 1:
		number, err = strconv.Atoi(firstTwoNumbers[:1])
		if err != nil {
			return 0, err
		}
		number *= 11
		return number, nil
	default:
		number, err = strconv.Atoi(firstTwoNumbers[:2])
		if err != nil {
			return 0, err
		}
		return number, nil
	}
}

func calculateSum(input *os.File) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		number, err := number(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	sum, err := calculateSum(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
