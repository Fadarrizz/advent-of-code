package main

import (
	_ "embed"
	"fmt"
	_ "fmt"
	"log"
	_ "os"
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
	var disk []int
	id := -1
	for i, block := range input {
		length, _ := strconv.Atoi(string(block))

		var v int
		if i % 2 == 0 {
			id++
			v = id
		} else {
			v = -1
		}

		for range length {
			disk = append(disk, v)
		}		
	}

	left := 0
	right := len(disk) - 1
	for left < right {
		if disk[left] != -1 {
			left++
			continue
		}

		if disk[right] == -1 {
			right--
			continue
		}

		tmp := disk[left]
		disk[left] = disk[right]
		disk[right] = tmp

		left++
		right--
	}

	total := 0
	for i, v := range disk {
		if v == -1 {
			break
		}

		total += i * v
	}

	return total, nil
}

func part2() (int, error) {
	var disk []int
	id := -1
	for i, block := range input {
		length, _ := strconv.Atoi(string(block))

		var v int
		if i % 2 == 0 {
			id++
			v = id
		} else {
			v = -1
		}

		for range length {
			disk = append(disk, v)
		}		
	}

	ll := 0
	lr := 0
	rl := len(disk) - 1
	rr := len(disk) - 1
	for rl > 0  {
		// Set ll and lr to start and end index of consecutive nums
		for disk[ll] != -1 && ll < rl {
			ll += 1
		}

		lr = ll + 1
		for disk[lr] == disk[ll] && lr < rl {
			lr += 1
		}

		// Set rl and rr to start and end index of consecutive nums
		for disk[rr] == -1 && rr > lr {
			rr -= 1
		}

		rl = rr - 1
		for disk[rl] == disk[rr] && rl > lr {
			rl -= 1
		}

		// No big enough empty space found. Move to next chunk (new R)
		if lr >= rl {
			ll = 0
			lr = 0

			rr = rl - 1
			rl = rr
			continue
		}

		l := lr - ll
		r := rr - rl

		// When R doesn't fit into L, search for new L
		if r > l {
			ll = lr + 1
			lr = ll
			continue
		}

		// R fits into L. Swap values
		for i := range r {
			tmp := disk[ll + i]
			disk[ll + i] = disk[rr - i]
			disk[rr - i] = tmp
		}

		// dump.P(ll, lr, rl, rr, l, r)

		ll = 0
		lr = 0

		rr = rl
		rl = rr

		// printDisk(disk)
	}

	total := 0
	for i, v := range disk {
		if v != -1 {
			total += i * v
		}
	}

	return total, nil
}

func lines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func printDisk(disk []int) {
	for _, v := range disk {
		x := strconv.Itoa(v)
		if x == "-1" {
			x = "."
		}
		fmt.Print(x)
	}
	fmt.Print("\n")
}
