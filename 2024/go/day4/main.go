package main

import (
	"advent-of-code/2024/go/coordinates"
	"advent-of-code/2024/go/grid"
	_ "embed"
	"errors"
	"log"
	_ "os"
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
	lines := lines()

	height := len(lines)
	width := len(lines[0])

	grid, err := grid.New[rune](height, width)
	if err != nil {
		return 0, errors.New("Invalid grid")
	}

	for i, line := range lines {
		for j, char := range line {
			grid.Set(coordinates.New(i, j), char)
		}
	}

	total := 0
	for row := range grid.Height {
		for col := range grid.Width {
			c := coordinates.New(row, col)
			l1, _ := grid.Get(c)

			if l1 != 'X' {
				continue
			}

			for _, dir := range coordinates.AllDirections() {
				c2 := coordinates.Position(c, dir)
				l2, err := grid.Get(c2)
				if err != nil || l2 != 'M' {
					continue;
				}

				c2 = coordinates.Position(c2, dir)
				l3, err := grid.Get(c2)
				if err != nil || l3 != 'A' {
					continue;
				}

				c2 = coordinates.Position(c2, dir)
				l4, err := grid.Get(c2)
				if err != nil || l4 != 'S' {
					continue;
				}

				total++
			}
		} 
	}

	return total, nil
}

func part2() (int, error) {
	lines := lines()

	height := len(lines)
	width := len(lines[0])

	grid, err := grid.New[rune](height, width)
	if err != nil {
		return 0, errors.New("Invalid grid")
	}

	for i, line := range lines {
		for j, char := range line {
			grid.Set(coordinates.New(i, j), char)
		}
	}

	total := 0
	for row := range grid.Height {
		for col := range grid.Width {
			c := coordinates.New(row, col)
			l, _ := grid.Get(c)

			if l != 'A' {
				continue
			}

			cd := coordinates.Position(c, coordinates.NORTHWEST)
			nw, err := grid.Get(cd)
			if err != nil || (nw != 'M' && nw != 'S') {
				continue;
			}

			cd = coordinates.Position(c, coordinates.NORTHEAST)
			ne, err := grid.Get(cd)
			if err != nil || (ne != 'M' && ne != 'S') {
				continue;
			}

			cd = coordinates.Position(c, coordinates.SOUTHEAST)
			se, err := grid.Get(cd)
			if err != nil || ((nw == 'M' && se != 'S') || (nw == 'S' && se != 'M')) {
				continue
			}

			cd = coordinates.Position(c, coordinates.SOUTHWEST)
			sw, err := grid.Get(cd)
			if err != nil || (ne == 'M' && sw != 'S') || (ne == 'S' && sw != 'M') {
				continue
			}
			
			total++
		} 
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
