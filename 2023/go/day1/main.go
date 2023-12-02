package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	part1, err := Part1()
	if err != nil {
		return
	}

	part2, err := Part2()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		left := strings.IndexAny(line, "0123456789")
		right := strings.LastIndexAny(line, "0123456789")

		calibration, err := strconv.Atoi(fmt.Sprintf("%c%c", line[left], line[right]))
		if err != nil {
			return -1, err
		}

		total += calibration
	}

	return total, nil
}

func Part2() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := []string{}

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				nums = append(nums, string(line[i]))
			}

			for key, num := range numbers {
				if strings.HasPrefix(line[i:], key) {
					nums = append(nums, num)
				}
			}
		}

		left := nums[0]
		right := nums[len(nums)-1]

		calibration, err := strconv.Atoi(fmt.Sprintf("%s%s", left, right))
		if err != nil {
			return -1, err
		}

		total += calibration
	}

	return total, nil
}
