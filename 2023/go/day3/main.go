package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	part1, err := part1()
	if err != nil {
		log.Fatal(err)
	}

	part2, err := part2()
	if err != nil {
		log.Fatal(err)
	}

	println(part1)
	println(part2)
}

func part1() (int, error) {
	lines := readlines()
	result := 0

	for row, line := range lines {
		var digits []rune
		for col, value := range line {
			// look for digit
			if unicode.IsDigit(value) {
				digits = append(digits, value)
			}

			// all digits have been collected
			if !unicode.IsDigit(value) || col == len(line) - 1 {
				if len(digits) > 0 {
					digitslen := len(digits)
					_, _, ok := findAround(
						len(lines),
						len(line),
						row,
						col - digitslen,
						digitslen,
						func (y int, x int) bool {
							return lines[y][x] != '.' && !unicode.IsNumber(rune(lines[y][x]))
						},
					)
					
					if ok {
						number, err := strconv.Atoi(string(digits))
						if err != nil {
							return 0, err
						}
						result += number
					}
				}

				digits = []rune{}
			}
		}
	}

	return result, nil
}

func part2() (int, error) {
	lines := readlines()
	result := 0

	var seenStars = make(map[int]int)
	for row, line := range lines {
		var digits []rune
		for col, value := range line {
			if unicode.IsDigit(value) {
				digits = append(digits, value)
			}

			if !unicode.IsDigit(value) || col == len(line) - 1 {
				if len(digits) > 0 {
					digitslen := len(digits)
					starRow, starCol, ok := findAround(
						len(lines),
						len(line),
						row,
						col - digitslen,
						digitslen,
						func (y int, x int) bool {
							return lines[y][x] == '*'
						},
					)
					
					if ok {
						number, err := strconv.Atoi(string(digits))
						if err != nil {
							return 0, err
						}
						
						index := calculateIndex(starRow, starCol, len(line))
						matchingValue, seen := seenStars[index] 
						if seen {
							result += number * matchingValue
						} else {
							seenStars[index] = number
						}
					}
				}

				digits = []rune{}
			}
		}
	}

	return result, nil
}

func findAround(width int, height int, y int, x int, offset int, comparefn func(int, int) bool) (int, int, bool) {
	for row := y - 1; row <= y + 1; row++ {
		for col := x - 1; col <= x + offset; col++ {
			if row >= 0 && row < width && col >= 0 && col < height {
				if comparefn(row, col) {
					return row, col, true
				}
			}
		}
	}

	return 0, 0, false
}

func calculateIndex(row int, col int, numCols int) int {
	return row * numCols + col
}

func readlines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
