package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color int

const (
	Red Color = iota
	Blue
	Green
)

var (
	colors = map[string]Color{
		"red":   Red,
		"blue":  Blue,
		"green": Green,
	}
)

type Cubes struct {
	N     int
	color Color
}

func parseCubes(s string) (*Cubes, error) {
	parts := strings.Split(s, " ")

	n, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	color, ok := colors[parts[1]]
    if ok == false {
        return nil, errors.New("unknown color encoutered.")
    }

	return &Cubes{n, color}, nil
}

type Set []*Cubes

func parseSet(s string) (*Set, error) {
	var set Set
	for _, cubesString := range strings.Split(s, ", ") {
		cubes, err := parseCubes(cubesString)

		if err != nil {
			return nil, err
		}

		set = append(set, cubes)
	}

	return &set, nil
}

type Game struct {
	id   int
	sets []*Set
}

func (g Game) highestCubeAmount(color Color) int {
	highestAmount := 0
	for _, set := range g.sets {
		for _, cubes := range *set {
			if cubes.color == color && cubes.N > highestAmount {
				highestAmount = cubes.N
			}
		}
	}

	return highestAmount
}

func parseGame(s string) (*Game, error) {
	parts := strings.Split(s, ": ")

	id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		return nil, err
	}

	var sets []*Set
	for _, subsetList := range strings.Split(parts[1], "; ") {
		set, err := parseSet(subsetList)
		if err != nil {
			return nil, err
		}

		sets = append(sets, set)
	}

    return &Game{id, sets}, nil
}

func main() {
	part1, err := part1()
	if err != nil {
		println(err)
	}

	println(part1)
}

func part1() (int, error) {
    var games []*Game
	for _, line := range readLines() {
        game, err := parseGame(line)
        if err != nil {
            return 0, err
        }

        games = append(games, game)
	}
	
	var result int
	for _, game := range games {
		reds := game.highestCubeAmount(Red)
		blues := game.highestCubeAmount(Blue)
		greens := game.highestCubeAmount(Green)

		if reds <= 12 && blues <= 14 && greens <= 13 {
			result += game.id
		}
	}

	return result, nil
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
