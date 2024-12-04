package grid

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alsm/aoc2024/aoc"
)

type Vector struct {
	Point     aoc.Point
	Direction aoc.Point
}

func (v Vector) String() string {
	return fmt.Sprintf("Point: %v  Direction: %v", v.Point, v.Direction)
}

var Directions4 = []aoc.Point{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

var Directions8 = []aoc.Point{
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: -1},
	{X: 0, Y: -1},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
}

type Grid[T any] struct {
	xLen      int64
	yLen      int64
	state     [][]T
	movements []aoc.Point
}

func New[T any](xLen, yLen int64, movements []aoc.Point) *Grid[T] {
	state := make([][]T, yLen)
	for y := int64(0); y < yLen; y++ {
		state[y] = make([]T, xLen)
	}
	return &Grid[T]{
		xLen:      xLen,
		yLen:      yLen,
		state:     state,
		movements: movements,
	}
}

func NewWithDefault[T any](xLen, yLen int64, movements []aoc.Point, def T) *Grid[T] {
	state := make([][]T, yLen)
	for y := int64(0); y < yLen; y++ {
		state[y] = make([]T, xLen)
		for i := range state[y] {
			state[y][i] = def
		}
	}
	return &Grid[T]{
		xLen:      xLen,
		yLen:      yLen,
		state:     state,
		movements: movements,
	}
}

func (g *Grid[T]) IsValid(x, y int64) bool {
	switch {
	case x < 0, x >= g.xLen, y < 0, y >= g.yLen:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) IsValidPoint(p aoc.Point) bool {
	switch {
	case p.X < 0, p.X >= g.xLen, p.Y < 0, p.Y >= g.yLen:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) Neighbours(p aoc.Point) []aoc.Point {
	var ret []aoc.Point

	for _, m := range g.movements {
		np := p.Add(m)
		if g.IsValidPoint(np) {
			ret = append(ret, np)
		}
	}

	return ret
}

func (g *Grid[T]) XLen() int64 {
	return g.xLen
}

func (g *Grid[T]) YLen() int64 {
	return g.yLen
}

func (g *Grid[T]) GetSliceToEdge(x, y int64, movement aoc.Point) []T {
	var ret []T
	p := aoc.Point{X: x, Y: y}
	for ; g.IsValidPoint(p); p = p.Add(movement) {
		ret = append(ret, g.GetState(p.Y, p.X))
	}

	return ret
}

func (g *Grid[T]) GetSliceInDirectionP(p aoc.Point, direction aoc.Point, count int) []T {
	var ret []T
	for ; g.IsValidPoint(p); p = p.Add(direction) {
		ret = append(ret, g.GetStateP(p))
		count -= 1
		if count == 0 {
			break
		}
	}

	return ret
}

func (g *Grid[T]) SetState(x, y int64, state T) {
	if g.IsValid(x, y) {
		g.state[y][x] = state
	}
}

func (g *Grid[T]) SetStateP(p aoc.Point, state T) {
	if g.IsValidPoint(p) {
		g.state[p.Y][p.X] = state
	}
}

func (g *Grid[T]) GetState(x, y int64) T {
	return g.state[y][x]
}

func (g *Grid[T]) GetStateP(p aoc.Point) T {
	return g.state[p.Y][p.X]
}

func (g *Grid[T]) StateString() string {
	var ret strings.Builder

	for _, y := range g.state {
		for _, x := range y {
			ret.WriteString(fmt.Sprintf("%v", x))
		}
		ret.WriteRune('\n')
	}

	return ret.String()
}

func (g *Grid[T]) StateStringInvertY() string {
	var ret strings.Builder

	for y := len(g.state) - 1; y >= 0; y-- {
		for _, x := range g.state[y] {
			ret.WriteString(fmt.Sprintf("%v", x))
		}
		ret.WriteRune('\n')
	}

	return ret.String()
}

func (g *Grid[T]) StateMap() map[aoc.Point]T {
	ret := make(map[aoc.Point]T)

	for y, l := range g.state {
		for x, s := range l {
			ret[aoc.Point{X: int64(x), Y: int64(y)}] = s
		}
	}

	return ret
}

func (g *Grid[T]) StateMapWhere(f func(T) bool) map[aoc.Point]T {
	ret := make(map[aoc.Point]T)

	for y, l := range g.state {
		for x, s := range l {
			if f(s) {
				ret[aoc.Point{X: int64(x), Y: int64(y)}] = s
			}
		}
	}

	return ret
}

func (g *Grid[T]) Clone() *Grid[T] {
	ng := Grid[T]{
		xLen:      g.xLen,
		yLen:      g.yLen,
		movements: slices.Clone(g.movements),
		state:     make([][]T, g.yLen),
	}

	for yi := range g.state {
		ng.state[yi] = slices.Clone(g.state[yi])
	}

	return &ng
}
