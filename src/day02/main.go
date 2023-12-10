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
	idPossibility := map[int]bool{}
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
		possibility := true
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
					if count > 12 {
						possibility = false
					}
					idPossibility[id] = possibility
				case "green":
					if count > 13 {
						possibility = false
					}
					idPossibility[id] = possibility
				case "blue":
					if count > 14 {
						possibility = false
					}
					idPossibility[id] = possibility
				default:
					continue
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
	sum := 0
	for id, possibility := range idPossibility {
		if possibility {
			sum += id
		}
	}
	fmt.Println(sum)
}
