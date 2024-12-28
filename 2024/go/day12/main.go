package main

import (
	"advent-of-code/2024/go/coordinates"
	"advent-of-code/2024/go/grid"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

type Region struct {
	plantType rune
	plants map[coordinates.Coordinate]bool
	perimeter int
}

func (r *Region) addPlant(c coordinates.Coordinate) {
	r.plants[c] = true
}

func (r *Region) incPerimeter() {
	r.perimeter++
}

func (r *Region) sides() int {
	cornersCandidates :=  make(map[coordinates.Corner]bool)
	for pos := range r.plants {
		for _, corner := range pos.Corners() {
			cornersCandidates[corner] = true
		}
	}

	sides := 0

	opposingCorners1 := []bool{true, false, true, false}
	opposingCorners2 := []bool{false, true, false, true}

	for cornerCandidate := range cornersCandidates {
		var surroundingSquares []bool
		squaresInRegion := 0
		for _, square := range cornerCandidate.SurroundingSquares() {
			_, exists := r.plants[square]

			if exists {
				squaresInRegion++
			}

			surroundingSquares = append(surroundingSquares, exists)
		}

		if squaresInRegion == 1 || squaresInRegion == 3 {
			sides++
		}

		if (surroundingSquares[0] == opposingCorners1[0] && surroundingSquares[1] == opposingCorners1[1] && surroundingSquares[2] == opposingCorners1[2] && surroundingSquares[3] == opposingCorners1[3]) || 
			(surroundingSquares[0] == opposingCorners2[0] && surroundingSquares[1] == opposingCorners2[1] && surroundingSquares[2] == opposingCorners2[2] && surroundingSquares[3] == opposingCorners2[3]) {
			sides += 2
		}
	} 

	return sides
}

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

	var regions []*Region
	var visited = make(map[coordinates.Coordinate]bool)
	walkFn := grid.Walk()

	var pos coordinates.Coordinate
	var walkErr error
	for true {
		pos, walkErr = walkFn()
		if walkErr != nil {
			break;
		}
		
		if _, ok := visited[pos]; !ok {
			region := createRegion(grid, pos, visited)
			regions = append(regions, region)
		}
	}

	total := 0
	for _, region := range regions {
		total += len(region.plants) * region.perimeter
	}

	return total, nil
}

func part2() (int, error) {
	lines := lines()
	grid := initGrid(lines)

	var regions []*Region
	var visited = make(map[coordinates.Coordinate]bool)
	walkFn := grid.Walk()

	var pos coordinates.Coordinate
	var walkErr error
	for true {
		pos, walkErr = walkFn()
		if walkErr != nil {
			break;
		}
		
		if _, ok := visited[pos]; !ok {
			region := createRegion(grid, pos, visited)
			regions = append(regions, region)
		}
	}

	total := 0
	for _, region := range regions {
		total += len(region.plants) * region.sides()
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func initGrid(s []string) *grid.Grid[rune] {
	var grid, _ = grid.New[rune](len(s), len(s[0]))
	for y, line := range s {
		for x, val := range line {
			grid.Set(coordinates.New(y, x), val)
		}
	}

	return grid
}

func createRegion(g *grid.Grid[rune], c coordinates.Coordinate, v map[coordinates.Coordinate]bool) *Region {
	val, _ := g.Get(c)
	region := &Region{plantType: val, plants: make(map[coordinates.Coordinate]bool)}

	dfs(g, c, v, region)

	return region
}

func dfs(g *grid.Grid[rune], c coordinates.Coordinate, v map[coordinates.Coordinate]bool, r *Region) {
	_, exists := v[c]
	val, _ := g.Get(c)
	if g.OutOfBounds(c) || exists || val != r.plantType {
		return
	}

	v[c] = true
	r.addPlant(c)

    for _, dir := range coordinates.Directions() {
		pos := coordinates.Position(c, dir)
		val, err := g.Get(pos)
		if err != nil || val != r.plantType {
			r.incPerimeter()
		} else {
			dfs(g, pos, v, r)
		}
	}
}
