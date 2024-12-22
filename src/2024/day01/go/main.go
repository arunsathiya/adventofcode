package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NumberPairs struct {
	a []int
	b []int
}

func readAndSortNumbers(input *os.File) NumberPairs {
	scanner := bufio.NewScanner(input)
	pairs := NumberPairs{
		a: make([]int, 0),
		b: make([]int, 0),
	}

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		firstN, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		secondN, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		pairs.a = append(pairs.a, firstN)
		pairs.b = append(pairs.b, secondN)
	}

	sort.Ints(pairs.a)
	sort.Ints(pairs.b)
	return pairs
}

func calculateDistances(pairs NumberPairs) int {
	sum := 0
	for i := 0; i < len(pairs.a); i++ {
		diff := math.Abs(float64(pairs.a[i] - pairs.b[i]))
		sum += int(diff)
	}
	return sum
}

func similarityScore(pairs NumberPairs) int {
	sum := 0
	for _, left := range pairs.a {
		count := 0
		for _, right := range pairs.b {
			if left == right {
				count++
			}
		}
		sum += left * count
	}
	return sum
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	pairs := readAndSortNumbers(input)
	distances := calculateDistances(pairs)
	similarity := similarityScore(pairs)
	fmt.Println(distances)
	fmt.Println(similarity)
}
