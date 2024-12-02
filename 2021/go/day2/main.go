package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

type Coordinate = [2]int;

//go:embed input.txt
var input string;

func main() {
	part1, err := part1();
	if err != nil {
		log.Fatal(err);
	}

	part2, err := part2();
	if err != nil {
		log.Fatal(err);
	}

	println(part1);
	println(part2);
}

func part1() (int, error) {
	pos := Coordinate{0,0};

	for _, line := range lines() {
		parts := strings.Split(line, " ");
		value, err := strconv.ParseInt(parts[1], 10, 0);

		if err != nil {
			log.Fatalf("Invalid int: %v\n", err);
		}

		switch parts[0] {
		case "forward":
			pos[0] += int(value);
		case "up":
			pos[1] -= int(value);
		case "down":
			pos[1] += int(value);
		}
	}
	
	return pos[0] * pos[1], nil
}

func part2() (int, error) {
	pos := Coordinate{0,0};
	aim := 0;

	for _, line := range lines() {
		parts := strings.Split(line, " ");
		value, err := strconv.ParseInt(parts[1], 10, 0);

		if err != nil {
			log.Fatalf("Invalid int: %v\n", err);
		}

		switch parts[0] {
		case "forward":
			pos[0] += int(value);
			pos[1] += aim * int(value);
		case "up":
			aim -= int(value);
		case "down":
			aim += int(value);
		}
	}
	
	return pos[0] * pos[1], nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n");
}
