package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day01Of2022(input *os.File) (int, error) {
	scanner := bufio.NewScanner(input)
	value := ""
	for scanner.Scan() {
		scanned := scanner.Text()
		value += scanned + "\n"
	}
	blocks := strings.Split(value, "\n\n")
	highestSum := 0
	for _, calories := range blocks {
		sum := 0
		calorie := strings.Split(calories, "\n")
		for i := 0; i < len(calorie); i++ {
			if calorie[i] != "" {
				calorie, err := strconv.Atoi(calorie[i])
				if err != nil {
					return 0, err
				}
				sum += calorie
			}
		}
		if sum > highestSum {
			highestSum = sum
		}
	}
	return highestSum, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	sum, err := Day01Of2022(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
