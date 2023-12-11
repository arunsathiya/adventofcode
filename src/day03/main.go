package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func numbers(input *os.File) (int, error) {
	re := regexp.MustCompile(`\d+`)
	lines := make([]string, 0)
	invalidNumbers := make([]int, 0)
	sum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		matchesInLine := re.FindAllString(line, -1)
		for _, match := range matchesInLine {
			number, err := strconv.Atoi(match)
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}
	}
	for lineIndex, line := range lines {
		if lineIndex != 0 && lineIndex != len(lines)-1 {
			matchesInLine := re.FindAllStringIndex(line, -1)
			for _, match := range matchesInLine {
				matchIndexAtLineLevelStart := match[0]
				matchIndexAtLineLevelEnd := match[1] - 1
				leftOk := (matchIndexAtLineLevelStart == 0 || string(line[matchIndexAtLineLevelStart-1]) == ".")
				rightOk := (matchIndexAtLineLevelEnd == len(line)-1 || string(line[matchIndexAtLineLevelEnd+1]) == ".")
				if leftOk && rightOk {
					for i := matchIndexAtLineLevelStart - 1; i <= matchIndexAtLineLevelEnd+1; i++ {
						if i >= 0 && i < len(line) {
							prevLine, nextLine := strings.Split(lines[lineIndex-1], ""), strings.Split(lines[lineIndex+1], "")
							if i < len(prevLine) && i < len(nextLine) && prevLine[i] == "." && nextLine[i] == "." {
								invalidNumber, err := strconv.Atoi(line[match[0]:match[1]])
								if err != nil {
									log.Fatal(err)
								}
								if len(invalidNumbers) == 0 || invalidNumbers[len(invalidNumbers)-1] != invalidNumber {
									invalidNumbers = append(invalidNumbers, invalidNumber)
								}
							}
						}
					}
				}
			}
		}
		if lineIndex == 0 {
			matchesInLine := re.FindAllStringIndex(line, -1)
			for _, match := range matchesInLine {
				matchIndexAtLineLevelStart := match[0]
				matchIndexAtLineLevelEnd := match[1] - 1
				leftOk := (matchIndexAtLineLevelStart == 0 || string(line[matchIndexAtLineLevelStart-1]) == ".")
				rightOk := (matchIndexAtLineLevelEnd == len(line)-1 || string(line[matchIndexAtLineLevelEnd+1]) == ".")
				if leftOk && rightOk {
					for i := matchIndexAtLineLevelStart - 1; i <= matchIndexAtLineLevelEnd+1; i++ {
						if i >= 0 && i < len(line) {
							nextLine := strings.Split(lines[lineIndex+1], "")
							if i < len(nextLine) && nextLine[i] == "." {
								invalidNumber, err := strconv.Atoi(line[match[0]:match[1]])
								if err != nil {
									log.Fatal(err)
								}
								if len(invalidNumbers) == 0 || invalidNumbers[len(invalidNumbers)-1] != invalidNumber {
									invalidNumbers = append(invalidNumbers, invalidNumber)
								}
							}
						}
					}
				}
			}
		}
		if lineIndex == len(lines)-1 {
			matchesInLine := re.FindAllStringIndex(line, -1)
			for _, match := range matchesInLine {
				matchIndexAtLineLevelStart := match[0]
				matchIndexAtLineLevelEnd := match[1] - 1
				leftOk := (matchIndexAtLineLevelStart == 0 || string(line[matchIndexAtLineLevelStart-1]) == ".")
				rightOk := (matchIndexAtLineLevelEnd == len(line)-1 || string(line[matchIndexAtLineLevelEnd+1]) == ".")
				if leftOk && rightOk {
					for i := matchIndexAtLineLevelStart - 1; i <= matchIndexAtLineLevelEnd+1; i++ {
						if i >= 0 && i < len(line) {
							prevLine := strings.Split(lines[lineIndex-1], "")
							if i < len(prevLine) && prevLine[i] == "." {
								invalidNumber, err := strconv.Atoi(line[match[0]:match[1]])
								if err != nil {
									log.Fatal(err)
								}
								if len(invalidNumbers) == 0 || invalidNumbers[len(invalidNumbers)-1] != invalidNumber {
									invalidNumbers = append(invalidNumbers, invalidNumber)
								}
							}
						}
					}
				}
			}
		}
	}
	for _, invalidNumber := range invalidNumbers {
		sum -= invalidNumber
	}
	return sum, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	sum, err := numbers(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
