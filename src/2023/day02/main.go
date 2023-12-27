package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type colorCount map[string]int

func identifyCubes(input *os.File) (map[int]colorCount, error) {
	idColor := make(map[int]colorCount)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		gameIdBlock := strings.Split(line, ":")[0]
		id, err := strconv.Atoi(strings.Split(gameIdBlock, " ")[1])
		if err != nil {
			return nil, err
		}
		cubesBlock := strings.Split(line, ":")[1]
		cubesBlock = cubesBlock[1:]
		attempts := strings.Split(cubesBlock, "; ")
		if idColor[id] == nil {
			idColor[id] = make(colorCount)
		}
		for _, attempt := range attempts {
			colors := strings.Split(attempt, ", ")
			for _, color := range colors {
				countColorCombo := strings.Split(color, " ")
				count, err := strconv.Atoi(countColorCombo[0])
				if err != nil {
					return nil, err
				}
				switch countColorCombo[1] {
				case "red":
					if count > idColor[id]["red"] {
						idColor[id]["red"] = count
					}
				case "green":
					if count > idColor[id]["green"] {
						idColor[id]["green"] = count
					}
				case "blue":
					if count > idColor[id]["blue"] {
						idColor[id]["blue"] = count
					}
				default:
					continue
				}
			}
		}
	}
	return idColor, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	cubes, err := identifyCubes(input)
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, colors := range cubes {
		sum += colors["red"] * colors["green"] * colors["blue"]
	}
	fmt.Println(sum)
}
