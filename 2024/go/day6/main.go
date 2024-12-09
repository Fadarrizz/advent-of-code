package main

import (
	"advent-of-code/2024/go/coordinates"
	"advent-of-code/2024/go/matrix"
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

	matrix := initMatrix(lines)

	startingPosition := matrix.Find('^')

	visited, _ := walk(matrix, startingPosition)

	return len(visited), nil
}

func part2() (int, error) {
	lines := lines()

	matrix := initMatrix(lines)

	startingPosition := matrix.Find('^')

	visited, _ := walk(matrix, startingPosition)

	total := 0
	for c := range visited {
		matrix.Set(c, '#')
		_, err := walk(matrix, startingPosition)
		if err != nil {
			total++
		}
		matrix.Set(c, '.')
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func initMatrix(s []string) matrix.Matrix[rune] {
	height := len(s)
	width := len(s[0])

	matrix, _ := matrix.New[rune](height, width)

	for i, line := range s {
		for j, char := range line {
			matrix.Set(coordinates.New(i, j), char)
		}
	}

	return *matrix
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

func walk(m matrix.Matrix[rune], c coordinates.Coordinate) (map[coordinates.Coordinate]bool, error) {
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
