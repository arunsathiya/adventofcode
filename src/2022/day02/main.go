package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func winCondition(opponent, me string) (int, error) {
	if (me == "X" && opponent == "C") || (me == "Y" && opponent == "A") || (me == "Z" && opponent == "B") {
		return 1, nil
	} else if (opponent == "A" && me == "Z") || (opponent == "B" && me == "X") || (opponent == "C" && me == "Y") {
		return -1, nil
	}
	return 0, nil
}

func Day02Of2022PartA(rounds []string) (int, error) {
	totalSum := 0
	for _, round := range rounds {
		sum := 0
		opponent := strings.Split(round, " ")[0]
		me := strings.Split(round, " ")[1]
		win, err := winCondition(opponent, me)
		if err != nil {
			return 0, err
		}
		switch win {
		case 1:
			sum += 6
		case 0:
			sum += 3
		case -1:
			sum += 0
		}
		switch me {
		case "X":
			sum += 1
		case "Y":
			sum += 2
		case "Z":
			sum += 3
		}
		totalSum += sum
	}
	return totalSum, nil
}

func Day02Of2022PartB(rounds []string) (int, error) {
	totalSum := 0
	for _, round := range rounds {
		sum := 0
		opponent := strings.Split(round, " ")[0]
		winCondition := strings.Split(round, " ")[1]
		switch winCondition {
		case "X":
			sum += 0
			switch opponent {
			case "A":
				sum += 3
			case "B":
				sum += 1
			case "C":
				sum += 2
			}
		case "Y":
			sum += 3
			switch opponent {
			case "A":
				sum += 1
			case "B":
				sum += 2
			case "C":
				sum += 3
			}
		case "Z":
			sum += 6
			switch opponent {
			case "A":
				sum += 2
			case "B":
				sum += 3
			case "C":
				sum += 1
			}
		}
		totalSum += sum
	}
	return totalSum, nil
}

func parseInput(input *os.File) ([]string, error) {
	scanner := bufio.NewScanner(input)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	rounds, err := parseInput(input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	sumOfRounds, err := Day02Of2022PartA(rounds)
	if err != nil {
		log.Fatalf(err.Error())
	}
	input.Seek(0, 0)
	partBSumOfRounds, err := Day02Of2022PartB(rounds)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(sumOfRounds)
	fmt.Println(partBSumOfRounds)
}
