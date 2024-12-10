package main

import (
	_ "embed"
	"fmt"
	_ "fmt"
	"log"
	_ "os"
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
	total := 0
	for _, line := range lines() {
		values := strings.Split(line, ": ")
		target, _ := strconv.Atoi(values[0])

		var nums []int
		for _, n := range strings.Split(strings.Trim(values[1], " "), " ") {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		if search(nums[0], target, nums[1:]) {
			total += target
		}
	}

	return total, nil
}

func part2() (int, error) {
	total := 0
	for _, line := range lines() {
		values := strings.Split(line, ": ")
		target, _ := strconv.Atoi(values[0])

		var nums []int
		for _, n := range strings.Split(strings.Trim(values[1], " "), " ") {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		if search(nums[0], target, nums[1:]) {
			total += target
		}
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func search(current int, target int, nums []int) bool {
	if current == target && len(nums) == 0 {
		return true
	} 

	if len(nums) == 0 || current > target {
		return false
	} 

	return search(pipe(current, nums[0]), target, nums[1:]) || 
		search(current * nums[0], target, nums[1:]) || 
		search(current + nums[0], target, nums[1:]) 
}

func pipe(a, b int) (int) {
	r, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))

	return r
}
