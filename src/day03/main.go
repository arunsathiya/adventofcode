package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func numbers(input *os.File) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	digit := regexp.MustCompile(`\d`)
	lines := make([]string, 0)
	invalid := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for lineIndex, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			matchIndicesAtLineLevel := make([]int, 0)
			for i := match[0]; i < match[1]; i++ {
				matchIndicesAtLineLevel = append(matchIndicesAtLineLevel, i)
			}
			for _, matchIndexAtLineLevel := range matchIndicesAtLineLevel {
				if lineIndex != 0 && lineIndex != len(lines)-1 {
					if matchIndexAtLineLevel != 0 && matchIndexAtLineLevel != len(strings.Split(line, ""))-1 {
						if strings.Split(lines[lineIndex-1], "")[matchIndexAtLineLevel] == "." && strings.Split(lines[lineIndex+1], "")[matchIndexAtLineLevel] == "." && (strings.Split(lines[lineIndex], "")[matchIndexAtLineLevel-1] == "." || digit.MatchString(strings.Split(lines[lineIndex], "")[matchIndexAtLineLevel-1])) && (strings.Split(lines[lineIndex], "")[matchIndexAtLineLevel+1] == "." || digit.MatchString(strings.Split(lines[lineIndex], "")[matchIndexAtLineLevel-1])) && strings.Split(lines[lineIndex-1], "")[matchIndexAtLineLevel-1] == "." && strings.Split(lines[lineIndex-1], "")[matchIndexAtLineLevel+1] == "." && strings.Split(lines[lineIndex+1], "")[matchIndexAtLineLevel-1] == "." && strings.Split(lines[lineIndex+1], "")[matchIndexAtLineLevel+1] == "." {
							invalidNumber, _ := strconv.Atoi(lines[lineIndex][match[0]:match[1]])
							invalid = append(invalid, invalidNumber)
						}
					}
				}
			}
		}
	}
	return nil, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	numbers, err := numbers(input)
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
}
