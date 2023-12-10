package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func identifyCubes(input *os.File) (map[int]bool, error) {
	idPossibility := map[int]bool{
		1: true,
		2: false,
		3: false,
		4: true,
	}
	scanner := bufio.NewScanner(input)
	possibility := true
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameIdBlock := strings.Split(line, ":")[0]
		id, _ = strconv.Atoi(strings.Split(gameIdBlock, " ")[1])
		cubesBlock := strings.Split(line, ":")[1]
		cubesBlock = cubesBlock[1:]
		attempts := strings.Split(cubesBlock, "; ")
		for _, attempt := range attempts {
			colors := strings.Split(attempt, ", ")
			for _, color := range colors {
				switch color {
				case "red":
					red, err := strconv.Atoi(strings.Split(color, " red")[0])
					if err != nil {
						return nil, err
					}
					if red > 12 {
						possibility = false
					}
					idPossibility[id] = possibility
				case "green":
					green, err := strconv.Atoi(strings.Split(color, " green")[0])
					if err != nil {
						return nil, err
					}
					if green > 13 {
						possibility = false
					}
					idPossibility[id] = possibility
				case "blue":
					blue, err := strconv.Atoi(strings.Split(color, " blue")[0])
					if err != nil {
						return nil, err
					}
					if blue > 14 {
						possibility = false
					}
					idPossibility[id] = possibility
				}
			}
		}
	}
	return idPossibility, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	idPossibility, err := identifyCubes(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(idPossibility)
}
