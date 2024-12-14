package main

import (
	"advent-of-code/2024/go/coordinates"
	"advent-of-code/2024/go/grid"
	_ "embed"
	_ "fmt"
	"log"
	// "os"
	"strconv"
	"strings"

	// "github.com/gookit/goutil/dump"
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
	grid := initGrid(lines)

	total := 0
	for row := range grid.Height {
		for col := range grid.Width {
			// Use map as a Set
			ends := make(map[coordinates.Coordinate]bool)

			pos := coordinates.New(row, col) 
			if v, _ := grid.Get(pos); v == 0 {
				findEnds(pos, grid, ends)

				total += len(ends)
			}
		}
	}

	return total, nil
}

func part2() (int, error) {
	lines := lines()
	grid := initGrid(lines)

	total := 0
	for row := range grid.Height {
		for col := range grid.Width {
			// Use slice to store all possibilities of trails
			var ends []coordinates.Coordinate

			pos := coordinates.New(row, col) 
			if v, _ := grid.Get(pos); v == 0 {
				findAll(pos, grid, &ends)

				total += len(ends)
			}
		}
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func initGrid(s []string) *grid.Grid[int] {
	grid, _ := grid.New[int](len(s), len(s[0]))

	for i, line := range s {
		for j, val := range line {
			n, _ := strconv.Atoi(string(val))

			grid.Set(coordinates.New(i, j), n)
		}
	}

	return grid
}

func findEnds(pos coordinates.Coordinate, grid *grid.Grid[int], ends map[coordinates.Coordinate]bool) {
	dirs := []coordinates.Direction{
		coordinates.NORTH,
		coordinates.EAST,
		coordinates.SOUTH,
		coordinates.WEST,
	}

	val, _ := grid.Get(pos)
	if val == 9 {
		ends[pos] = true
		return
	}

	for _, dir := range dirs {
		next := coordinates.Position(pos, dir)

		nextVal, _ := grid.Get(next)
		if nextVal == val + 1 {
			findEnds(next, grid, ends)
		}
	}
}

func findAll(pos coordinates.Coordinate, grid *grid.Grid[int], ends *[]coordinates.Coordinate) {
	dirs := []coordinates.Direction{
		coordinates.NORTH,
		coordinates.EAST,
		coordinates.SOUTH,
		coordinates.WEST,
	}

	val, _ := grid.Get(pos)
	if val == 9 {
		*ends = append(*ends, pos)
		return
	}

	for _, dir := range dirs {
		next := coordinates.Position(pos, dir)

		nextVal, _ := grid.Get(next)
		if nextVal == val + 1 {
			findAll(next, grid, ends)
		}
	}
}
