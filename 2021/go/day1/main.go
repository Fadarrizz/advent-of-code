package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed example.txt
var input string

func main() {
	part1, err := part1();
	if err != nil {
		log.Fatal(err)
	}

	part2, err := part2();
	if err != nil {
		log.Fatal(err)
	}

	println(part1)
	println(part2)
}

func part1() (int, error) {
	increases := 0

	lines := lines()
	for i := 1; i < len(lines); i++ {
		if lines[i - 1] < lines[i] {
			increases += 1
		}
	}

	return increases, nil
}

func part2() (int, error) {
	increases := 0

	lines := lines()
	for i := 5; i < len(lines); i += 3 {
		a := lines[i - 5] + lines[i - 4] + lines[i - 3]
		b := lines[i - 2] + lines[i - 1] + lines[i]

		if a < b {
			increases += 1
		}
	}

	return increases, nil
}

func Map[T any, U any](vals []T, fn func(T) (U, error)) ([]U, error) {
	us := make([]U, len(vals))
	for i, v := range vals {
		u, err := fn(v)
		if err != nil {
			return nil, err
		}
		us[i] = u
	}
	return us, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
