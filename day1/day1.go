package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"github.com/alsm/aoc2024/aoc"
)

func readInput(b io.Reader) ([]int, []int) {
	ls := bufio.NewScanner(b)

	var left, right []int
	var l, r int
	for ls.Scan() {
		fmt.Sscanf(ls.Text(), "%d   %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func partOne(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	z := aoc.Zip(left, right)
	diffs := aoc.Map(z, func(v []int) int {
		return aoc.Abs(v[0] - v[1])
	})
	return aoc.Sum(diffs)
}

func partTwo(left, right []int) int {
	count := aoc.Tally(right)

	return aoc.Reduce(left, 0, func(s int, v int) int {
		return s + v*count[v]
	})
}

func main() {
	data, err := os.ReadFile("day1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	left, right := readInput(bytes.NewReader(data))
	log.Println("Part 1:", partOne(left, right))
	log.Println("Part 2:", partTwo(left, right))
}
