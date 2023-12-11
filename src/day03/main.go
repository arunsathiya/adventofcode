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

func numbers(input *os.File) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	lines := make([]string, 0)
	invalidNumbers := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
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
	}
	fmt.Println(invalidNumbers)
	return nil, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	_, err = numbers(input)
	if err != nil {
		log.Fatal(err)
	}
}
