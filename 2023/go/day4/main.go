package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1, err := part1()
	if err != nil {
		println(err)
	}

	part2, err := part2()
	if err != nil {
		log.Fatal(err)
	}

	println(part1)
	println(part2)
}

func part1() (int, error) {
	result := 0
	for _, line := range readLines() {
		winning, owned := parseLine(line)

		points := 0
		for _, n := range owned {
			if slices.Contains(winning, n) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		result += points
	}

	return result, nil
}

func part2() (int, error) {
	lines := readLines()

	cards := []int{}
	for i := 0; i < len(lines); i++ {
		cards = append(cards, 1)
	}

	for i, line := range lines {
		winning, owned := parseLine(line)

		points := 0
		for _, n := range owned {
			if slices.Contains(winning, n) {
				points++
			}
		}

		for k := 0; k < cards[i]; k++ {
			for j := 1; j <= points; j++ {
				cards[i+j] += 1
			}
		}	
	}

	sum := 0
	for i := range cards {
		sum += cards[i]
	}

	return sum, nil
}

func parseNumbers(str string) []int {
	var result []int
	for _, part := range strings.Fields(str) {
		if num, err := strconv.Atoi(part); err == nil {
			result = append(result, num)
		}
	}

	return result
}

func parseLine(str string) ([]int, []int) {
	sets := strings.Split(strings.SplitAfter(str, ":")[1], "|")

	winning := parseNumbers(sets[0])
	owned := parseNumbers(sets[1])

	return winning, owned
}

func readLines() []string {
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
