package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/alsm/aoc2024/aoc"
)

func readInput(r io.Reader) map[int64][]int64 {
	ret := make(map[int64][]int64)
	d, _ := io.ReadAll(r)
	for line := range slices.Values(strings.Split(string(d), "\n")) {
		parts := strings.Split(line, ": ")
		ret[aoc.Atoi64(parts[0])] = aoc.Map(strings.Split(parts[1], " "), func(s string) int64 {
			return aoc.Atoi64(s)
		})
	}

	return ret
}

func generateOperators(x int, n int) [][]int {
	var ret [][]int
	t := make([]int, n)
	for {
		ret = append(ret, slices.Clone(t))
		i := n - 1
		for ; i >= 0 && t[i] == x-1; i-- {
			t[i] = 0
		}
		if i < 0 {
			break
		}
		t[i]++
	}
	return ret
}

func calculate(nums []int64, ops []int) int64 {
	if len(nums) == 1 {
		return nums[0]
	}
	s := nums[0]
	for i, v := range nums[1:] {
		switch ops[i] {
		case 0:
			s *= v
		case 1:
			s += v
		case 2:
			s = aoc.Atoi64(fmt.Sprintf("%d%d", s, v))
		}
	}
	return s
}

func partOne(ops map[int64][]int64) int64 {
	valid := aoc.SelectMap(ops, func(v int64, n []int64) bool {
		o := generateOperators(2, len(n)-1)
		vs := aoc.Map(o, func(x []int) int64 {
			return calculate(n, x)
		})
		return aoc.Any(vs, func(x int64) bool {
			return x == v
		})
	})

	return aoc.Sum(slices.Collect(maps.Keys(valid)))
}

func partTwo(ops map[int64][]int64) int64 {
	valid := aoc.SelectMap(ops, func(v int64, n []int64) bool {
		o := generateOperators(3, len(n)-1)
		vs := aoc.Map(o, func(x []int) int64 {
			return calculate(n, x)
		})
		return aoc.Any(vs, func(x int64) bool {
			return x == v
		})
	})

	return aoc.Sum(slices.Collect(maps.Keys(valid)))
}

func main() {
	f, _ := os.ReadFile("day7.txt")

	ops := readInput(bytes.NewReader(f))
	log.Println("Part 1:", partOne(ops))
	log.Println("Part 2:", partTwo(ops))
}
