package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(input *os.File) ([]string, error) {
	scanner := bufio.NewScanner(input)
	masterPairs := make([]string, 0)
	for scanner.Scan() {
		masterPairs = append(masterPairs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return masterPairs, nil
}

func Day04Of2022PartA(pairs []string) (int, error) {
	fullyContained := 0
	for _, pair := range pairs {
		ranges := strings.Split(pair, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")
		range1Start, _ := strconv.Atoi(range1[0])
		range1End, _ := strconv.Atoi(range1[1])
		range2Start, _ := strconv.Atoi(range2[0])
		range2End, _ := strconv.Atoi(range2[1])
		if (range1Start <= range2Start && range1End >= range2End) || (range2Start <= range1Start && range2End >= range1End) {
			fullyContained++
		}
	}
	return fullyContained, nil
}

func Day04Of2022PartB(pairs []string) (int, error) {
	partiallyContained := 0
	for _, pair := range pairs {
		ranges := strings.Split(pair, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")
		range1Start, _ := strconv.Atoi(range1[0])
		range1End, _ := strconv.Atoi(range1[1])
		range2Start, _ := strconv.Atoi(range2[0])
		range2End, _ := strconv.Atoi(range2[1])
		if (range1Start <= range2Start && range1End >= range2Start) || (range2Start <= range1Start && range2End >= range1Start) {
			partiallyContained++
		}
	}
	return partiallyContained, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	pairs, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fullyContained, err := Day04Of2022PartA(pairs)
	if err != nil {
		log.Fatal(err)
	}
	partiallyContained, err := Day04Of2022PartB(pairs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullyContained)
	fmt.Println(partiallyContained)
}
