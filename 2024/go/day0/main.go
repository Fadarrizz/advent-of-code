package main

import (
	_"embed"
	_"os"
	_"fmt"
	"log"
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
	return 0, nil
}

func part2() (int, error) {
	return 0, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
