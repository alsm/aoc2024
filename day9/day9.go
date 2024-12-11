package main

import (
	"iter"
	"log"
	"maps"
	"os"
	"slices"

	"github.com/alsm/aoc2024/aoc"
)

func expand(in [][2]int) []int {
	var b []int
	for i, v := range in {
		b = append(b, aoc.Repeat(i, v[0])...)
		b = append(b, aoc.Repeat(-1, v[1])...)
	}
	return b
}

func compact(a []int) []int {
	bptr := len(a) - 1
	next := func(x, y int) int {
		for ; y > x; y-- {
			if a[y] != -1 {
				return y
			}
		}
		return 0
	}
	for i, v := range a {
		if v == -1 {
			bptr = next(i, bptr)
			if bptr == 0 {
				break
			}
			a[i], a[bptr] = a[bptr], a[i]
		}
	}

	return a
}

func Countdown(x int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := x; i >= 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

type run struct {
	i int
	l int
}

func genMap(a []int) map[int]run {
	ret := make(map[int]run)
	for n, c := range aoc.Tally(a) {
		i := slices.Index(a, n)
		ret[n] = run{
			i: i,
			l: c,
		}
	}

	return ret
}

func defrag(a []int) []int {
	findGap := func(i int, s int) int {
		for ; i < len(a); i++ {
			if a[i] == -1 {
				j := i
				for ; j < len(a) && a[j] == -1; j++ {
				}
				if j-i >= s {
					return i
				}
			}
		}
		return 0
	}

	indices := make([]int, 10)
	numLocs := genMap(a)
	delete(numLocs, -1)
	x := slices.Collect(maps.Keys(numLocs))
	slices.Sort(x)
	slices.Reverse(x)
	for n := range slices.Values(x) {
		r := numLocs[n]
		i := findGap(indices[r.l], r.l)
		if i != 0 && i < r.i {
			for z := range r.l {
				a[i+z] = n
				a[r.i+z] = -1
			}
			indices[r.l] = i + r.l
		}
	}

	return a
}

func parseInput(d string) [][2]int {
	if len(d)%2 != 0 {
		d += "0"
	}
	parts := aoc.Chunk([]rune(d), 2)
	return aoc.Map(parts, func(v []rune) [2]int {
		return [2]int{aoc.Atoi(string(v[0])), aoc.Atoi(string(v[1]))}
	})
}

func partOne(d string) int64 {
	x := expand(parseInput(d))
	x = compact(x)
	return aoc.Sum(aoc.MapWithIndex(x, func(i int, v int) int64 {
		if v == -1 {
			return 0
		}
		return int64(i * v)
	}))
}

func partTwo(d string) int64 {
	x := expand(parseInput(d))
	x = defrag(x)
	return aoc.Sum(aoc.MapWithIndex(x, func(i int, v int) int64 {
		if v == -1 {
			return 0
		}
		return int64(i * v)
	}))
}

func main() {
	d, _ := os.ReadFile("day9.txt")

	log.Println("Part 1:", partOne(string(d)))
	log.Println("Part 2:", partTwo(string(d)))
}
