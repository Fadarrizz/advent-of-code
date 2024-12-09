package main

import (
	_ "embed"
	_ "fmt"
	"log"
	_ "os"
	"slices"
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

type Rule struct {
	A int
	B int
}

func part1() (int, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := rules(parts[0])
	updates := updates(parts[1])

	total := 0
	for _, update := range updates {
		correct := true
		for _, rule := range rules {
			leftIdx := slices.Index(update, rule.A)
			rightIdx := slices.Index(update, rule.B)

			if leftIdx == -1 || rightIdx == -1 {
				continue
			}

			if leftIdx > rightIdx {
				correct = false
				break;
			}
		}

		if correct {
			middleIdx := (len(update) - 1) / 2
			total += update[middleIdx]
		}
	}

	return total, nil
}

func part2() (int, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := rules(parts[0])
	updates := updates(parts[1])

	incorrectUpdates := make(map[int]bool)
	for i, update := range updates {
		correct := 0
		for correct != len(rules) {
			correct = 0
			for _, rule := range rules {
				leftIdx := slices.Index(update, rule.A)
				rightIdx := slices.Index(update, rule.B)

				if leftIdx == -1 || rightIdx == -1 || leftIdx < rightIdx {
					correct++
					continue
				}

				update[leftIdx] = rule.B
				update[rightIdx] = rule.A

				incorrectUpdates[i] = true
			}
		}
	}

	total := 0
	for i := range incorrectUpdates {
		update := updates[i]	
		middleIdx := (len(update) - 1) / 2
		total += update[middleIdx]
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func rules(s string) []Rule {
	ruleLines := strings.Split(s, "\n")

	rules := make([]Rule, len(ruleLines))
	for i, line := range ruleLines {
		r := strings.Split(line, "|")

		a, _ := strconv.Atoi(r[0])
		b, _ := strconv.Atoi(r[1])

		rule := Rule{
			A: a,
			B: b,
		}

		rules[i] = rule
	}

	return rules
}

func updates(s string) [][]int {
	updateLines := strings.Split(s, "\n")

	updates := make([][]int, len(updateLines))
	for i, update := range updateLines {
		pageLines := strings.Split(update, ",")
		pages := make([]int, len(pageLines))	
		for j, page := range pageLines {
			p, _ := strconv.Atoi(page)
			pages[j] = p
		}
		updates[i] = pages
	}

	return updates
}
