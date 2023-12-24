package main

import (
	"advent-of-code/2023/go/pkg/conv"
	"embed"
	"log"
	"strings"
)

//go:embed example.txt
var input string

//go:embed example*.txt
var examples embed.FS

const MAX_INT = int(^uint(0) >> 1)

type Rule struct {
	destination int
	source		int
	length		int
}

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
	seeds := parseSeeds(lines[0])
	rules := make([][]*Rule, 7)
	index := -1

	for _, line := range lines[2:] {
		line := strings.TrimSpace(line)

		if strings.Contains(line, ":") {
			index++;
			continue
		}

		rule := parseRule(line)
		if rule != nil {
			rules[index] = append(rules[index], rule)
		}
	}

	result := MAX_INT
	for _, seed := range seeds {
		result = min(resolve(rules, seed), result)
	}

	return result, nil
}

func part2() (int, error) {
	lines := lines()
	seedRanges := parseSeeds(lines[0])
	rules := make([][]*Rule, 7)
	index := -1

	for _, line := range lines[2:] {
		line := strings.TrimSpace(line)

		if strings.Contains(line, ":") {
			index++;
			continue
		}

		rule := parseRule(line)
		if rule != nil {
			rules[index] = append(rules[index], rule)
		}
	}

	result := MAX_INT
	index = 0
	for index < len(seedRanges) {
		seedLength := seedRanges[index + 1]

		for i := 0; i < seedLength; i++ {
			seed := seedRanges[index] + seedLength

			result = min(resolve(rules, seed), result)
		}

		index += 2
	}

	return result, nil
}

func resolve(rules [][]*Rule, seed int) int {
	n := seed
	for _, ranges := range rules {
		tmp := n
		for _, rule := range ranges {
			if n >= rule.source && n < rule.source + rule.length {
				tmp = rule.destination + (n - rule.source)
			}
		}
		n = tmp
	}

	return n
}

func parseRule(s string) *Rule {
	parts := conv.StrsToInts(strings.Fields(s))

	if len(parts) != 3 {
		return nil
	}

	return &Rule{
		destination: parts[0],
		source: parts[1],
		length: parts[2],
	}
}

func parseSeeds(s string) []int {
	index := strings.IndexAny(s, ":") + 1
	seedStrings := strings.Fields(s[index:])

	return conv.StrsToInts(seedStrings)
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
