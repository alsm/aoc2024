package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alsm/aoc2024/aoc"
)

func readInput(r io.Reader) [][]int {
	ls := bufio.NewScanner(r)

	var lines [][]int
	for ls.Scan() {
		numstr := strings.Split(ls.Text(), " ")
		nums := aoc.Map(func(v string) int {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			return n
		}, slices.Values(numstr))
		lines = append(lines, slices.Collect(nums))
	}

	return lines
}

func partOne(l [][]int) int {
	filtered := aoc.Filter(func(l []int) bool {
		z := aoc.Zip(slices.Values(l[:len(l)-1]), slices.Values(l[1:]))
		inc := aoc.All(func(v aoc.Zipped[int, int]) bool {
			diff := v.V1 - v.V2
			return diff >= 1 && diff <= 3
		}, z)
		dec := aoc.All(func(v aoc.Zipped[int, int]) bool {
			diff := v.V1 - v.V2
			return diff <= -1 && diff >= -3
		}, z)
		return inc || dec
	}, slices.Values(l))

	return len(slices.Collect(filtered))
}

func partTwo(l [][]int) int {
	filtered := aoc.Filter(func(l []int) bool {
		var variants [][]int
		for i := range len(l) {
			variants = append(variants, slices.Delete(slices.Clone(l), i, i+1))
		}
		return aoc.Any(func(sl []int) bool {
			z := aoc.Zip(slices.Values(sl[:len(sl)-1]), slices.Values(sl[1:]))
			inc := aoc.All(func(v aoc.Zipped[int, int]) bool {
				diff := v.V1 - v.V2
				return diff >= 1 && diff <= 3
			}, z)
			dec := aoc.All(func(v aoc.Zipped[int, int]) bool {
				diff := v.V1 - v.V2
				return diff <= -1 && diff >= -3
			}, z)
			return inc || dec
		}, slices.Values(variants))
	}, slices.Values(l))

	return len(slices.Collect(filtered))
}

func main() {
	data, err := os.ReadFile("day2.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := readInput(bytes.NewReader(data))
	log.Println("Part 1:", partOne(lines))
	log.Println("Part 2:", partTwo(lines))
}
