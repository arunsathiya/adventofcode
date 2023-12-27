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
	calories := strings.Split(value, "\n\n")
	highestSum := 0
	for _, perPersonCalories := range calories {
		sum := 0
		calorie := strings.Split(perPersonCalories, "\n")
		for i := 0; i < len(calorie); i++ {
			if calorie[i] != "" {
				calorie, err := strconv.Atoi(calorie[i])
				if err != nil {
					return 0, err
				}
				sum += calorie
			}
		}
		highestSum = sum
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
