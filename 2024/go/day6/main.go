package main

import (
	"advent-of-code/2024/go/coordinates"
	"advent-of-code/2024/go/grid"
	_ "embed"
	"errors"
	"fmt"
	"log"
	_"os"
	"slices"
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

	grid := initgrid(lines)

	startingPosition := grid.Find('^')

	visited, _ := walk(grid, startingPosition)

	return len(visited), nil
}

func part2() (int, error) {
	lines := lines()

	grid := initgrid(lines)

	startingPosition := grid.Find('^')

	visited, _ := walk(grid, startingPosition)

	total := 0
	for c := range visited {
		grid.Set(c, '#')
		_, err := walk(grid, startingPosition)
		if err != nil {
			total++
		}
		grid.Set(c, '.')
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func initgrid(s []string) grid.Grid[rune] {
	height := len(s)
	width := len(s[0])

	grid, _ := grid.New[rune](height, width)

	for i, line := range s {
		for j, char := range line {
			grid.Set(coordinates.New(i, j), char)
		}
	}

	return *grid
}

func turn(d coordinates.Direction) coordinates.Direction {
	dirs := []coordinates.Direction{
		coordinates.NORTH,
		coordinates.EAST,
		coordinates.SOUTH,
		coordinates.WEST,
	}

	idx := slices.Index(dirs, d)

	return dirs[(idx + 1) % len(dirs)]
}

func walk(m grid.Grid[rune], c coordinates.Coordinate) (map[coordinates.Coordinate]bool, error) {
	d := coordinates.NORTH

	visited := make(map[coordinates.Coordinate]bool)
	visited[c] = true

	duplicates := make(map[string]bool)
	duplicates[formatPositionWithDirection(c, d)] = true

	for m.InBounds(c) {
		next := coordinates.Position(c, d)
		val, _ := m.Get(next)

		switch val {
		// Turn when in front of #
		case '#':
			d = turn(d)
		// Walk in direction otherwise
		default:
			c = next
			visited[c] = true
		}

		if duplicates[formatPositionWithDirection(c, d)] {
			return nil, errors.New("Loop encountered")
		}

		duplicates[formatPositionWithDirection(c, d)] = true
	}

	return visited, nil
}

func formatPosition(c coordinates.Coordinate) string {
	return fmt.Sprint(c.X, c.Y)
}

func formatPositionWithDirection(c coordinates.Coordinate, d coordinates.Direction) string {
	return fmt.Sprint(c.X, c.Y, d.Name())
}
