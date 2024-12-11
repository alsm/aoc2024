package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alsm/aoc2024/aoc"
)

var cache = make(map[[2]int]int)

func blink(i int, blinks int) int {
	if blinks == 0 {
		return 1
	}
	if x, ok := cache[[2]int{i, blinks}]; ok {
		return x
	}
	is := strconv.Itoa(i)
	var x int
	switch {
	case i == 0:
		x = blink(1, blinks-1)
	case len(is)%2 == 0:
		x = blink(aoc.Atoi(is[:len(is)/2]), blinks-1) + blink(aoc.Atoi(is[len(is)/2:]), blinks-1)
	default:
		x = blink(i*2024, blinks-1)
	}
	cache[[2]int{i, blinks}] = x
	return x
}

func partOne(n []int) int {
	x := aoc.Map(n, func(i int) int {
		return blink(i, 25)
	})

	return aoc.Sum(x)
}

func partTwo(n []int) int {
	x := aoc.Map(n, func(i int) int {
		return blink(i, 75)
	})

	return aoc.Sum(x)
}

func main() {
	d, _ := os.ReadFile("day11.txt")
	n := aoc.Map(strings.Split(string(d), " "), func(s string) int {
		return aoc.Atoi(s)
	})

	log.Println("Part 1:", partOne(n))
	log.Println("Part 2:", partTwo(n))
}
