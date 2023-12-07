package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibration := make([]int, 0)
		for _, v := range line {
			v, err := strconv.Atoi(string(v))
			if len(calibration) == 2 {
				break
			}
			if err == nil {
				calibration = append(calibration, v)
			}
		}
		switch len(calibration) {
		case 1:
			sum += calibration[0]*10 + calibration[0]
		case 2:
			sum += calibration[0]*10 + calibration[1]
		default:
			sum += 0
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
