package main

import (
	"bytes"
	"io"
	"log"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/alsm/aoc2024/aoc"
	"github.com/alsm/aoc2024/aoc/grid"
)

func readInput(b io.Reader) *grid.Grid[rune] {
	d, err := io.ReadAll(b)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(d), "\n")

	g := grid.New[rune](int64(len(lines[0])), int64(len(lines)), grid.Directions8)
	for y, l := range lines {
		for x, v := range l {
			g.SetState(int64(x), int64(y), v)
		}
	}

	return g
}

func findXmas(g *grid.Grid[rune], p aoc.Point) int {
	maybes := aoc.Map(func(d aoc.Point) []rune {
		return g.GetSliceInDirectionP(p, d, 4)
	}, slices.Values(grid.Directions8))
	xmass := aoc.Filter(func(x []rune) bool {
		return string(x) == "XMAS"
	}, maybes)

	return len(slices.Collect(xmass))
}

func partOne(g *grid.Grid[rune]) int {
	xs := g.StateMapWhere(func(r rune) bool {
		return r == 'X'
	})

	count := aoc.Reduce(func(s int, v int) int {
		return s + v
	}, 0, aoc.Map(func(p aoc.Point) int {
		return findXmas(g, p)
	}, maps.Keys(xs)))

	return count
}

func findMas(g *grid.Grid[rune], p aoc.Point) bool {
	m1 := g.GetSliceInDirectionP(p.Add(aoc.Point{-1, -1}), aoc.Point{1, 1}, 3)
	m2 := g.GetSliceInDirectionP(p.Add(aoc.Point{1, -1}), aoc.Point{-1, 1}, 3)

	slices.Sort(m1)
	slices.Sort(m2)

	return string(m1) == "AMS" && string(m2) == "AMS"
}

func partTwo(g *grid.Grid[rune]) int {
	as := aoc.Filter(func(p aoc.Point) bool {
		return p.X != 0 && p.Y != 0 && p.X != g.XLen()-1 && p.Y != g.YLen()-1
	}, maps.Keys(g.StateMapWhere(func(r rune) bool {
		return r == 'A'
	})))

	return aoc.Count(func(p aoc.Point) bool {
		return findMas(g, p)
	}, as)
}

func main() {
	data, err := os.ReadFile("day4.txt")
	if err != nil {
		log.Fatalln(err)
	}

	g := readInput(bytes.NewReader(data))
	log.Println("Part 1:", partOne(g))
	log.Println("Part 2:", partTwo(g))
}
