package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day01Of2022PartA(blocks []string) (int, error) {
	highestSum := 0
	for _, block := range blocks {
		sum, err := processBlock(block)
		if err != nil {
			return 0, err
		}
		if sum > highestSum {
			highestSum = sum
		}
	}
	return highestSum, nil
}

func Day01Of2022PartB(blocks []string) (int, error) {
	return 0, nil
}

func parseInput(input *os.File) ([]string, error) {
	scanner := bufio.NewScanner(input)
	value := ""
	for scanner.Scan() {
		scanned := scanner.Text()
		value += scanned + "\n"
	}
	return strings.Split(value, "\n\n"), scanner.Err()
}

func processBlock(block string) (int, error) {
	sum := 0
	calories := strings.Split(block, "\n")
	for _, calorie := range calories {
		if calorie != "" {
			value, err := strconv.Atoi(calorie)
			if err != nil {
				return 0, err
			}
			sum += value
		}
	}
	return sum, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	blocks, err := parseInput(input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	highestSum, err := Day01Of2022PartA(blocks)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(highestSum)
}
