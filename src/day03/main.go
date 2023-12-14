package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

func numbers(input *os.File) (int, error) {
	re := regexp.MustCompile(`\d+`)
	lines := make([]string, 0)
	collectedForSum := make([]int, 0)
	sum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for lineIndex, line := range lines {
		matchesInLine := re.FindAllStringIndex(line, -1)
		for _, match := range matchesInLine {
			number, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				return 0, err
			}
			if nearSymbol(lines, lineIndex, match[0], match[1]) {
				collectedForSum = append(collectedForSum, number)
				sum += number
			}
		}
	}
	fmt.Println(collectedForSum)
	return sum, nil
}

func nearSymbol(lines []string, lineIndex, startIndex, endIndex int) bool {
	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		for col := startIndex; col <= endIndex; col++ {
			newRow, newCol := lineIndex+dir.dx, col+dir.dy
			if newRow < 0 || newRow >= len(lines) || newCol < 0 || newCol >= len(lines[newRow]) {
				continue
			}
			char := lines[newRow][newCol]
			if char != '.' && !isDigit(rune(char)) {
				return true
			}
		}
	}
	return false
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
