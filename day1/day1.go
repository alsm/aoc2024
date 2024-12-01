package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/alsm/aoc2024/aoc"
)

func readInput(b io.Reader) ([]int, []int) {
	ls := bufio.NewScanner(b)
	var lines []string
	for ls.Scan() {
		lines = append(lines, strings.TrimSpace(ls.Text()))
	}

	var left, right []int
	var l, r int
	for line := range slices.Values(lines) {
		fmt.Sscanf(line, "%d   %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func partOne(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	z := aoc.Zip(slices.Values(left), slices.Values(right))
	diffs := aoc.Map(func(v aoc.Zipped[int, int]) int {
		return aoc.Abs(v.V1 - v.V2)
	}, z)
	return aoc.Reduce(func(s int, v int) int {
		return s + v
	}, 0, diffs)
}

func partTwo(left, right []int) int {
	count := aoc.Tally(right)

	return aoc.Reduce(func(s int, v int) int {
		return s + v*count[v]
	}, 0, slices.Values(left))
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
