package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"slices"

	"github.com/alsm/aoc2024/aoc"
)

var muls = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var allincs = regexp.MustCompile(`(mul\((\d+),(\d+)\)|don't\(\)|do\(\))`)

func readInput(b io.Reader, r *regexp.Regexp) [][2]int {
	var ret [][2]int
	data, err := io.ReadAll(b)
	if err != nil {
		log.Fatalln(err)
	}

	do := true

	m := r.FindAllStringSubmatch(string(data), -1)
	for i := range slices.Values(m) {
		switch {
		case i[0] == "do()":
			do = true
		case i[0] == "don't()":
			do = false
		default:
			if do {
				ret = append(ret, [2]int{aoc.Atoi(i[len(i)-2]), aoc.Atoi(i[len(i)-1])})
			}
		}
	}

	return ret
}

func solve(ins [][2]int) int {
	return aoc.Reduce(func(s int, v [2]int) int {
		return s + v[0]*v[1]
	}, 0, slices.Values(ins))
}

func main() {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		log.Fatalln(err)
	}

	i1 := readInput(bytes.NewReader(data), muls)
	log.Println("Part 1:", solve(i1))
	i2 := readInput(bytes.NewReader(data), allincs)
	log.Println("Part 2:", solve(i2))
}
