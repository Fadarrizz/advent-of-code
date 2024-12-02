package main

import (
	_ "embed"
	_ "fmt"
	"log"
	_ "os"
	"sort"
	"strconv"
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
	left := make([]int, 1000);
	right := make([]int, 1000);

	for i, line := range lines() {
		s := strings.Split(line, " ");

		a, err := strconv.Atoi(s[0]);
		if err != nil {
			return -1, err
		}

		b, err := strconv.Atoi(s[3]);
		if err != nil {
			return -1, err
		}

		left[i] = a;
		right[i] = b;
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j];
	});

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j];
	});

	dists := make([]int, 1000);
	dist := 0;
	for i := 0; i < 1000; i++ {
		if left[i] < right[i] {
			dist = right[i] - left[i]
		} else {
			dist = left[i] - right[i]
		}
		dists[i] = dist
	}

	total := 0;
	for _, val := range dists {
		total += val;
	}

	return total, nil
}

func part2() (int, error) {

	lookup := map[int]int{}
	right := make([]int, 1000);
	for i, line := range lines() {
		s := strings.Split(line, " ");

		a, err := strconv.Atoi(s[0]);
		if err != nil {
			return -1, err
		}

		b, err := strconv.Atoi(s[3]);
		if err != nil {
			return -1, err
		}

		lookup[a] = 0;
		right[i] = b;
	}

	total := 0;
	for _, num := range right {
		if _, ok := lookup[num]; ok {
			total += num;
		}
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
