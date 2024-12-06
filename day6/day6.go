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

func readInput(b io.Reader) (aoc.Point, *grid.Grid[rune]) {
	var start aoc.Point
	d, _ := io.ReadAll(b)
	lines := strings.Split(strings.TrimSpace(string(d)), "\n")
	g := grid.New[rune](int64(len(lines[0])), int64(len(lines)), grid.Directions4)
	for y, l := range lines {
		for x, r := range l {
			g.SetState(int64(x), int64(y), r)
			if r == '^' {
				start.X = int64(x)
				start.Y = int64(y)
			}
		}
	}

	return start, g
}

type Guard struct {
	location  aoc.Point
	direction aoc.Point
}

func (g *Guard) Next() aoc.Point {
	return g.location.Add(g.direction)
}

func (g *Guard) Step() aoc.Point {
	g.location = g.location.Add(g.direction)
	return g.location
}

func (g *Guard) Turn() {
	g.direction = aoc.Point{
		X: -g.direction.Y,
		Y: g.direction.X,
	}
}

func getPath(s aoc.Point, g *grid.Grid[rune]) []aoc.Point {
	guard := Guard{
		location:  s,
		direction: aoc.Point{X: 0, Y: -1},
	}

	points := make(map[aoc.Point]struct{})

	g.SetStateP(guard.location, 'X')

	for g.IsValidPoint(guard.Next()) {
		if g.GetStateP(guard.Next()) == '#' {
			guard.Turn()
			continue
		}
		points[guard.Step()] = struct{}{}
	}

	return slices.Collect(maps.Keys(points))
}

func partOne(s aoc.Point, g *grid.Grid[rune]) int {
	return len(getPath(s, g))
}

type LocVec struct {
	PX, PY, DX, DY int64
}

func partTwo(s aoc.Point, g *grid.Grid[rune]) int {
	var loopCount int
	pointCount := make(map[LocVec]struct{})
	locs := getPath(s, g)
	locs = slices.DeleteFunc(locs, func(p aoc.Point) bool {
		return p == s
	})
	for p := range slices.Values(locs) {
		// for _, p := range []aoc.Point{{6, 7}} {
		g.SetStateP(p, '#')
		guard := Guard{
			location:  s,
			direction: aoc.Point{X: 0, Y: -1},
		}
		clear(pointCount)
		for g.IsValidPoint(guard.Next()) {
			if g.GetStateP(guard.Next()) == '#' {
				guard.Turn()
				continue
			}
			np := guard.Step()
			vec := LocVec{np.X, np.Y, guard.direction.X, guard.direction.Y}
			if _, ok := pointCount[vec]; ok {
				loopCount++
				break
			}
			pointCount[vec] = struct{}{}
		}
		g.SetStateP(p, '.')
	}

	return loopCount
}

func main() {
	data, _ := os.ReadFile("day6.txt")

	s, g := readInput(bytes.NewReader(data))
	log.Println("Part 1:", partOne(s, g.Clone()))
	log.Println("Part 2:", partTwo(s, g.Clone()))
}
