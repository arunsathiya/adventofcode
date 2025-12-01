[![Go Checks](https://github.com/arunsathiya/adventofcode/actions/workflows/go-tests.yml/badge.svg)](https://github.com/arunsathiya/adventofcode/actions/workflows/go-tests.yml)
[![Python Checks](https://github.com/arunsathiya/adventofcode/actions/workflows/python-tests.yml/badge.svg)](https://github.com/arunsathiya/adventofcode/actions/workflows/python-tests.yml)

# Advent of Code

My Advent of Code solutions in Go and Python.

## Quick Start (2025+)

Create a new day:
```bash
./scripts/new-day.sh 2025 01
```

This creates:
```
src/2025/day01/
├── input.txt    # Paste your puzzle input here
└── go/
    └── main.go  # Solution template
```

## How to Run

### Go
```bash
cd src/2025/day01/go
go run main.go
```

### Python (2022-2023)
```bash
cd src/2022/day01/python
python main.py
```

## Structure

```
src/
├── 2022/       # Go + Python solutions
├── 2023/       # Go + Python solutions
├── 2024/       # Go solutions
└── 2025/       # New year
```
