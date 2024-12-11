package main

import (
	"advent-of-code/2024/go/coordinates"
	"advent-of-code/2024/go/grid"
	_ "embed"
	_ "fmt"
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

	grid := initGrid(lines)

	lookup := make(map[rune][]coordinates.Coordinate)
	for i, line := range lines {
		for j, val := range line {
			if val != '.' {
				lookup[val] = append(lookup[val], coordinates.New(i, j))
			}
		}
	}

	total := 0
	for _, coords := range lookup {
		for i := 0; i < len(coords); i++  {
			for j := i + 1; j < len(coords); j++ {
				line := coordinates.Line{
					A: coords[i],
					B: coords[j],
				}
				extended := line.Extend(line.Length())

				val, _ := grid.Get(extended.A)
				if grid.InBounds(extended.A) && val != '#' {
					grid.Set(extended.A, '#')
					total++
				}

				val, _ = grid.Get(extended.B)
				if grid.InBounds(extended.B) && val != '#' {
					grid.Set(extended.B, '#')
					total++
				}
			}
		}
	}


	return total, nil
}

func part2() (int, error) {
	lines := lines()

	grid := initGrid(lines)

	lookup := make(map[rune][]coordinates.Coordinate)
	for i, line := range lines {
		for j, val := range line {
			if val != '.' {
				lookup[val] = append(lookup[val], coordinates.New(i, j))
			}
		}
	}

	for _, coords := range lookup {
		for i := 0; i < len(coords); i++  {
			for j := i + 1; j < len(coords); j++ {
				a := coords[i]
				b := coords[j]

				length := coordinates.Line{A: a, B: b}.Length()
				for true {
					extended := coordinates.Line{
						A: a,
						B: b,
					}.Extend(length)

					a = extended.A
					b = extended.B

					if !grid.InBounds(a) && !grid.InBounds(b) {
						break;
					}

					val, _ := grid.Get(a)
					if grid.InBounds(a) && val != '#' {
						grid.Set(a, '#')
					}

					val, _ = grid.Get(b)
					if grid.InBounds(b) && val != '#' {
						grid.Set(b, '#')
					}
				}
			}
		}
	}

	total := 0
	for i := range grid.Height {
		for j := range grid.Width {
			val, _ := grid.Get(coordinates.New(i, j))
			if val != '.' {
				total++
			}
		}
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func initGrid(s []string) *grid.Grid[rune] {
	grid, _ := grid.New[rune](len(s), len(s[0]))

	for i, line := range s {
		for j, val := range line {
			grid.Set(coordinates.New(i, j), val)
		}
	}

	return grid
}
