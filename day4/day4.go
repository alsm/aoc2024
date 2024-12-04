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
	ms := aoc.Map(func(np aoc.Point) grid.Vector {
		return grid.Vector{
			Point:     np,
			Direction: np.Sub(p),
		}
	}, aoc.Filter(func(p aoc.Point) bool { return g.GetState(p.X, p.Y) == 'M' }, slices.Values(g.Neighbours(p))))
	as := aoc.Map(func(v grid.Vector) grid.Vector {
		return grid.Vector{
			Point:     v.Point.Add(v.Direction),
			Direction: v.Direction,
		}
	}, aoc.Filter(func(v grid.Vector) bool {
		np := v.Point.Add(v.Direction)
		return g.IsValidPoint(np) && g.GetState(np.X, np.Y) == 'A'
	}, ms))
	ss := aoc.Map(func(v grid.Vector) grid.Vector {
		return grid.Vector{
			Point:     v.Point.Add(v.Direction),
			Direction: v.Direction,
		}
	}, aoc.Filter(func(v grid.Vector) bool {
		np := v.Point.Add(v.Direction)
		return g.IsValidPoint(np) && g.GetState(np.X, np.Y) == 'S'
	}, as))

	return len(slices.Collect(ss))
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
	x := []aoc.Point{{X: -1, Y: -1}, {X: -1, Y: 1}, {X: 1, Y: -1}, {X: 1, Y: 1}}
	a := (g.GetStateP(p.Add(x[0])) - g.GetStateP(p.Add(x[1]))) +
		(g.GetStateP(p.Add(x[2])) - g.GetStateP(p.Add(x[3])))
	b := (g.GetStateP(p.Add(x[0])) - g.GetStateP(p.Add(x[2]))) +
		(g.GetStateP(p.Add(x[1])) - g.GetStateP(p.Add(x[3])))

	return (aoc.Abs(a) == 12 || aoc.Abs(b) == 12) && !slices.Contains(slices.Collect(aoc.Map(func(np aoc.Point) rune {
		return g.GetStateP(p.Add(np))
	}, slices.Values(x))), 'A')
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
