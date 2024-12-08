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

func readInput(b io.Reader) *grid.Grid[string] {
	d, _ := io.ReadAll(b)
	lines := strings.Split(string(d), "\n")
	g := grid.New[string](int64(len(lines[0])), int64(len(lines)), grid.Directions8)
	for y, l := range lines {
		for x, r := range l {
			g.SetState(int64(x), int64(y), string(r))
		}
	}

	return g
}

func findAntiNodes(p1, p2 aoc.Point) (aoc.Point, aoc.Point) {
	diff := p1.Sub(p2)
	return p1.Add(diff), p2.Sub(diff)
}

func partOne(g *grid.Grid[string]) int {
	groups := make(map[string][]aoc.Point)
	aoc.EachMap(g.StateMapWhere(func(r string) bool {
		return r != "."
	}), func(p aoc.Point, c string) {
		groups[c] = append(groups[c], p)
	})

	aoc.Each(slices.Collect(maps.Values(groups)), func(ps []aoc.Point) {
		aoc.Each(aoc.Combinations(ps), func(pair [2]aoc.Point) {
			an1, an2 := findAntiNodes(pair[0], pair[1])
			g.SetStateP(an1, "#")
			g.SetStateP(an2, "#")
		})
	})

	return len(g.StateMapWhere(func(r string) bool {
		return r == "#"
	}))
}

func findAllAntiNodes(g *grid.Grid[string], p1, p2 aoc.Point) []aoc.Point {
	var ret []aoc.Point
	diff := p1.Sub(p2)
	for np := p1; g.IsValidPoint(np); np = np.Add(diff) {
		ret = append(ret, np)
	}
	for np := p2; g.IsValidPoint(np); np = np.Sub(diff) {
		ret = append(ret, np)
	}

	return ret
}

func partTwo(g *grid.Grid[string]) int {
	groups := make(map[string][]aoc.Point)
	aoc.EachMap(g.StateMapWhere(func(r string) bool {
		return r != "."
	}), func(p aoc.Point, c string) {
		groups[c] = append(groups[c], p)
	})

	aoc.Each(slices.Collect(maps.Values(groups)), func(ps []aoc.Point) {
		aoc.Each(aoc.Combinations(ps), func(pair [2]aoc.Point) {
			aoc.Each(findAllAntiNodes(g, pair[0], pair[1]), func(p aoc.Point) {
				g.SetStateP(p, "#")
			})
		})
	})

	return len(g.StateMapWhere(func(r string) bool {
		return r == "#"
	}))
}

func main() {
	data, _ := os.ReadFile("day8.txt")

	g := readInput(bytes.NewReader(data))
	log.Println("Part 1:", partOne(g.Clone()))
	log.Println("Part 2:", partTwo(g.Clone()))
}
