package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func parseInput(input *os.File) ([]string, error) {
	scanner := bufio.NewScanner(input)
	rucksacks := make([]string, 0)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}
	return rucksacks, nil
}

func Day03Of2022PartA(rucksacks []string) (int, error) {
	priorities := 0
	for _, rucksack := range rucksacks {
		items := strings.Split(rucksack, "")
		itemsA, itemsB := halves(items)
		sort.Slice(itemsA, func(i, j int) bool {
			return itemsA[i] < itemsA[j]
		})
		sort.Slice(itemsB, func(i, j int) bool {
			return itemsB[i] < itemsB[j]
		})
		for _, item := range itemsB {
			status, target, err := search(itemsA, item)
			if err != nil {
				return 0, err
			}
			if status {
				runeItem := []rune(target)[0]
				priorities += charToNumber(runeItem)
				break
			}
		}
	}
	return priorities, nil
}

func charToNumber(ch rune) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch - 'a' + 1)
	} else if ch >= 'A' && ch <= 'Z' {
		return int(ch - 'A' + 27)
	}
	return -1
}

func search(items []string, target string) (bool, string, error) {
	left, right := 0, len(items)-1
	for left <= right {
		mid := left + (right-left)/2
		if items[mid] == target {
			return true, target, nil
		}
		if items[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false, "", nil
}

func halves(items []string) ([]string, []string) {
	mid := len(items) / 2
	return items[:mid], items[mid:]
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	rucksacks, err := parseInput(input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	priorities, err := Day03Of2022PartA(rucksacks)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(priorities)
}
