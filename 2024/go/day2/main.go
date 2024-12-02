package main

import (
	_ "embed"
	_ "fmt"
	"log"
	_ "os"
	"strconv"
	"strings"
	_ "github.com/gookit/goutil/dump"
)

//go:embed input.txt
var input string

const INC = 1
const NEUTRAL = 0
const DEC = -1

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
	total := 0;
	for _, line := range lines() {
		levels, err := parseLevels(line)
		if err != nil {
			return 1, nil
		}

		if isSafeReport(levels) {
			total++
		}
	}

	return total, nil
}

func part2() (int, error) {
	total := 0;
	for _, line := range lines() {
		levels, err := parseLevels(line)
		if err != nil {
			return 1, nil
		}

		if isSafeReportWithDampening(levels) {
			total++
		}
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func parseLevels(report string) ([]int, error) {
	stringLevels := strings.Split(report, " ")

	levels := make([]int, len(stringLevels));
	for i, level := range stringLevels {
		num, err := strconv.Atoi(level)
		if err != nil {
			panic("Cannot convert level to int: $ERR")
		}

		levels[i] = num;
	}

	return levels, nil
}

func isSafeReport(levels []int) bool {
	dir := dir(levels[0], levels[1])
	if dir == NEUTRAL {
		return false;
	}

	for i := range levels {
		if i == 0 {
			continue
		}

		a := levels[i-1]
		b := levels[i]

		diff := 0;
		if dir == INC {
			diff = b - a;
		} else {
			diff = a - b;
		}

		if diff < 1 || diff > 3 {
			return false;
		}
	}

	return true;
}

func isSafeReportWithDampening(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	for i := range levels {
		levelsWithDampening := make([]int, 0)
		levelsWithDampening = append(levelsWithDampening, levels[:i]...)
		levelsWithDampening = append(levelsWithDampening, levels[i+1:]...)

		if isSafeReport(levelsWithDampening) {
			return true;
		}
	}

	return false;
}

func dir(a int, b int) int {
	if a < b {
		return INC
	} else if a > b {
		return DEC
	} else {
		return NEUTRAL
	}
}
