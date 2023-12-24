package main

import (
	"advent-of-code/2023/go/pkg/conv"
	_ "advent-of-code/2023/go/pkg/utils"
	_ "embed"
	_ "fmt"
	"log"
	_ "os"
	"strings"
)

//go:embed input.txt
var input string

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
	lines := lines()

	times := parseTimes(lines[0])
	distances := parseDistances(lines[1])

	beatProduct := 1
	for raceIdx := 0; raceIdx < len(times); raceIdx++ {
		time := times[raceIdx]
		distance := distances[raceIdx]

		beatCount := 0
		for timeIdx := 0; timeIdx <= time; timeIdx++ {
			product := timeIdx * (time - timeIdx)

			if product > distance {
				beatCount++
			}
		}

		beatProduct *= beatCount
	}

	return beatProduct, nil
}

func part2() (int, error) {
	lines := lines()

	time := parseInput(lines[0])
	distance := parseInput(lines[1])

	beatCount := 0
	for timeIdx := 0; timeIdx <= time; timeIdx++ {
		product := timeIdx * (time - timeIdx)

		if product > distance {
			beatCount++
		}
	}

	return beatCount, nil
}

func parseTimes(s string) []int {
	index := strings.Index(s, ":") + 1
	timeStrings := strings.Fields(s[index:])

	return conv.StrsToInts(timeStrings)
}

func parseDistances(s string) []int {
	index := strings.IndexAny(s, ":") + 1
	distanceStrings := strings.Fields(s[index:])

	return conv.StrsToInts(distanceStrings)
}

func parseInput(s string) int {
	return conv.StrToInt(strings.Replace(strings.Split(s, ": ")[1], " ", "", -1))
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
