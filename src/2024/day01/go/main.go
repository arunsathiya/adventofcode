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

func calculateDistances(input *os.File) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	a := make([]int, 0)
	b := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		first := strings.Split(line, "   ")[0]
		firstN, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}
		second := strings.Split(line, "   ")[1]
		secondN, err := strconv.Atoi(second)
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, firstN)
		b = append(b, secondN)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i := 0; i < len(a); i++ {
		diff := math.Abs(float64(a[i] - b[i]))
		sum += int(diff)
	}
	return sum, nil
}

func similarityScore(input *os.File) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	a := make([]int, 0)
	b := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		first := strings.Split(line, "   ")[0]
		firstN, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}
		second := strings.Split(line, "   ")[1]
		secondN, err := strconv.Atoi(second)
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, firstN)
		b = append(b, secondN)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for _, left := range a {
		count := 0
		for _, right := range b {
			if left == right {
				count++
			}
		}
		sum += left * count
	}
	return sum, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	distances, err := calculateDistances(input)
	if err != nil {
		log.Fatal(err)
	}
	input.Seek(0, 0)
	similarityScore, err := similarityScore(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(distances)
	fmt.Println(similarityScore)
}
