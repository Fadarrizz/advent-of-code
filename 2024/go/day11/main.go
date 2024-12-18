package main

import (
	_ "embed"
	"log"
	// "os"
	"strconv"
	"strings"

	// "github.com/gookit/goutil/dump"
)

//go:embed example.txt
var input string

type Result struct {
	num int
	depth int
}

func main() {
	part1, err := part1()
	if err != nil {
		log.Fatal(err)
	}
	
	println(part1)

	part2, err := part2()
	if err != nil {
		log.Fatal(err)
	}

	println(part2)
}

func part1() (int, error) {
	total := 0
	cache := make(map[*Result]int)
	for _, v := range strings.Split(strings.Trim(input, "\n"), " ") {
		n, _ := strconv.Atoi(v)

		total += blink(&Result{n, 25}, cache)
	}

	return total, nil
}

func part2() (int, error) {
	total := 0
	cache := make(map[*Result]int)
	for _, v := range strings.Split(strings.Trim(input, "\n"), " ") {
		n, _ := strconv.Atoi(v)

		total += blink(&Result{n, 75}, cache)
	}

	return total, nil
}

func blink(result *Result, cache map[*Result]int) int {
	num := result.num
	depth := result.depth

	if depth == 0 {
		return 1
	}

	if v, ok := cache[result]; ok {
		return v
	}

	if num == 0 {
		return blink(&Result{num, depth - 1}, cache)
	} 

	s := strconv.Itoa(num)
	if len(s) % 2 == 0 {
		mid := len(s) / 2

		n1, _ := strconv.Atoi(s[:mid])
		n2, _ := strconv.Atoi(s[mid:])

		return blink(&Result{n1, depth - 1}, cache) + blink(&Result{n2, depth - 1}, cache)
	} 

	return blink(&Result{num * 2024, depth - 1}, cache)
}
