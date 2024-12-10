package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	"github.com/alsm/aoc2024/aoc"
	"github.com/alsm/aoc2024/aoc/grid"
	"github.com/alsm/aoc2024/aoc/queue"
	"github.com/alsm/aoc2024/aoc/set"
)

func readInput(b io.Reader) *grid.Grid[int] {
	d, _ := io.ReadAll(b)
	lines := strings.Split(strings.TrimSpace(string(d)), "\n")
	g := grid.New[int](int64(len(lines[0])), int64(len(lines)), grid.Directions4)
	for y, l := range lines {
		for x, n := range l {
			g.SetState(int64(x), int64(y), aoc.Atoi(string(n)))
		}
	}

	return g
}

func findPaths(g *grid.Grid[int], p aoc.Point) int {
	paths := set.New[aoc.Point]()
	q := queue.Queue[aoc.Point]{p}

	for !q.Empty() {
		x := q.Get()
		if s := g.GetStateP(x); s == 8 {
			aoc.Each(aoc.Select(g.Neighbours(x), func(y aoc.Point) bool {
				return g.GetStateP(y) == 9
			}), func(y aoc.Point) {
				paths.Add(y)
			})
		} else {
			aoc.Each(aoc.Select(g.Neighbours(x), func(z aoc.Point) bool {
				return g.GetStateP(z) == s+1
			}), func(z aoc.Point) {
				q.Put(z)
			})
		}
	}

	return len(paths)
}

func findAllPaths(g *grid.Grid[int], p aoc.Point) int {
	var paths int
	q := queue.Queue[aoc.Point]{p}

	for !q.Empty() {
		x := q.Get()
		if s := g.GetStateP(x); s == 8 {
			paths += aoc.Count(g.Neighbours(x), func(y aoc.Point) bool {
				return g.GetStateP(y) == 9
			})
		} else {
			aoc.Each(aoc.Select(g.Neighbours(x), func(z aoc.Point) bool {
				return g.GetStateP(z) == s+1
			}), func(z aoc.Point) {
				q.Put(z)
			})
		}
	}

	return paths
}

func partOne(g *grid.Grid[int]) int {
	starts := g.StatesWhere(func(i int) bool {
		return i == 0
	})

	return aoc.Reduce(starts, 0, func(sum int, p aoc.Point) int {
		return sum + findPaths(g, p)
	})
}

func partTwo(g *grid.Grid[int]) int {
	starts := g.StatesWhere(func(i int) bool {
		return i == 0
	})

	return aoc.Reduce(starts, 0, func(sum int, p aoc.Point) int {
		return sum + findAllPaths(g, p)
	})
}

func main() {
	data, _ := os.ReadFile("day10.txt")

	g := readInput(bytes.NewReader(data))
	log.Println("Part 1:", partOne(g.Clone()))
	log.Println("Part 2:", partTwo(g.Clone()))
}
