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
			if search(itemsA, item) {
				runeItem := []rune(item)[0]
				priorities += charToNumber(runeItem)
				break
			}
		}
	}
	return priorities, nil
}

func Day03Of2022PartB(rucksacks []string) (int, error) {
	var groups [][]string
	priorities := 0
	for i := 0; i < len(rucksacks); i += 3 {
		end := i + 3
		if end > len(rucksacks) {
			end = len(rucksacks)
		}
		groups = append(groups, rucksacks[i:end])
	}
	for _, group := range groups {
		itemsA, itemsB, itemsC := group[0], group[1], group[2]
		arrA, arrB, arrC := strings.Split(itemsA, ""), strings.Split(itemsB, ""), strings.Split(itemsC, "")
		sort.Slice(arrA, func(i, j int) bool {
			return arrA[i] < arrA[j]
		})
		sort.Slice(arrB, func(i, j int) bool {
			return arrB[i] < arrB[j]
		})
		sort.Slice(arrC, func(i, j int) bool {
			return arrC[i] < arrC[j]
		})
		shortest, second, third := shortest(arrA, arrB, arrC)
		for _, item := range shortest {
			if search(second, item) && search(third, item) {
				priorities += charToNumber([]rune(item)[0])
				break
			}
		}
	}
	return priorities, nil
}

func shortest(arr1, arr2, arr3 []string) ([]string, []string, []string) {
	if len(arr1) <= len(arr2) && len(arr1) <= len(arr3) {
		return arr1, arr2, arr3
	}
	if len(arr2) <= len(arr1) && len(arr2) <= len(arr3) {
		return arr2, arr1, arr3
	}
	return arr3, arr1, arr2
}

func charToNumber(ch rune) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch - 'a' + 1)
	} else if ch >= 'A' && ch <= 'Z' {
		return int(ch - 'A' + 27)
	}
	return -1
}

func search(items []string, target string) bool {
	left, right := 0, len(items)-1
	for left <= right {
		mid := left + (right-left)/2
		if items[mid] == target {
			return true
		}
		if items[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
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
	input.Seek(0, 0)
	prioritiesPartB, err := Day03Of2022PartB(rucksacks)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(priorities)
	fmt.Println(prioritiesPartB)
}
