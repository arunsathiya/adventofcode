#!/bin/bash
set -e

YEAR=${1:-2025}
DAY=$2

if [ -z "$DAY" ]; then
    echo "Usage: ./scripts/new-day.sh [year] <day>"
    echo "Example: ./scripts/new-day.sh 2025 01"
    exit 1
fi

# Pad day with zero if needed
DAY=$(printf "%02d" "$DAY")

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BASE_DIR="$(dirname "$SCRIPT_DIR")"
DAY_DIR="$BASE_DIR/src/$YEAR/day$DAY"

if [ -d "$DAY_DIR" ]; then
    echo "Directory $DAY_DIR already exists"
    exit 1
fi

echo "Creating Advent of Code $YEAR Day $DAY..."

# Create directory structure
mkdir -p "$DAY_DIR/go"

# Create shared input file
touch "$DAY_DIR/input.txt"

# Create Go main file from template
if [ -f "$SCRIPT_DIR/templates/main.go.tmpl" ]; then
    cp "$SCRIPT_DIR/templates/main.go.tmpl" "$DAY_DIR/go/main.go"
else
    cat > "$DAY_DIR/go/main.go" << 'EOF'
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partA(lines []string) int {
	// TODO: Implement part A
	return 0
}

func partB(lines []string) int {
	// TODO: Implement part B
	return 0
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part A:", partA(lines))
	fmt.Println("Part B:", partB(lines))
}
EOF
fi

echo "Created:"
echo "  $DAY_DIR/input.txt"
echo "  $DAY_DIR/go/main.go"
echo ""
echo "Next steps:"
echo "  1. Paste your puzzle input into $DAY_DIR/input.txt"
echo "  2. cd $DAY_DIR/go && go run main.go"
