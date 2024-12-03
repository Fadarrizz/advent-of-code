package main

import (
	_ "embed"
	_ "fmt"
	"log"
	_ "os"
	"regexp"
	"strconv"
	"strings"
	_ "github.com/gookit/goutil/dump"
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
	total := 0
	for _, match := range matches(input) {
		x, y := parseMul(match)
		total += x * y
	}

	return total, nil
}

func part2() (int, error) {
	total := 0
	for _, match := range matchesWithDosDonts(input) {
		x, y := parseMul(match)
		total += x * y
	}

	return total, nil
}

func matches(s string) []string {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	return r.FindAllString(s, -1)
}

func matchesWithDosDonts(s string) []string {
	r := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\(\d+,\d+\))`)

	matches := r.FindAllString(s, -1)

	var result []string
	enabled := true

	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else {
			if enabled {
				result = append(result, match)
			}
		}
	}

	return result
}

func parseMul(s string) (int, int) {
	startIdx := 4
	endIdx := len(s) - 1
	commaIdx := strings.Index(s, ",")

	x, err := strconv.Atoi(s[startIdx:commaIdx])
	if err != nil {
		panic("")
	}

	y, err := strconv.Atoi(s[commaIdx+1:endIdx])
	if err != nil {
		panic("")
	}

	return x, y
}
