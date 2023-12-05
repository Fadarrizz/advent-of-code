package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	part1, err := part1()
	if err != nil {
		println(err)
	}

	part2, err := part2()
	if err != nil {
		println(err)
	}

	println(part1)
	println(part2)
}

func part1() (int, error) {
}

func part2() (int, error) {
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
