package aoc

import (
	"iter"
	"log"
	"slices"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Cycle[T any](s iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			for v := range s {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func LCM(nums []int) int {
	r := nums[0]
	for i := range slices.Values(nums[1:]) {
		r = (r * i) / GCD(r, i)
	}

	return r
}

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Sign[T constraints.Integer | constraints.Float](x T) T {
	switch {
	case x < 0:
		return -T(1)
	case x > 0:
		return 1
	default:
		return 0
	}
}

type Point struct {
	X int64
	Y int64
}

func (p *Point) MDistance(b Point) int64 {
	return Abs(p.X-b.X) + Abs(p.Y-b.Y)
}

func (p *Point) MDistanceXY(x, y int64) int64 {
	return Abs(p.X-x) + Abs(p.Y-y)
}

func (p *Point) Add(b Point) Point {
	return Point{
		X: p.X + b.X,
		Y: p.Y + b.Y,
	}
}

func (p *Point) Sub(b Point) Point {
	return Point{
		X: p.X - b.X,
		Y: p.Y - b.Y,
	}
}

func (p *Point) Line(b Point) []Point {
	var ret []Point
	dx := b.X - p.X
	dy := b.Y - p.Y
	sx := Sign(dx)
	sy := Sign(dy)

	num := Abs(dx + dy)
	for i := int64(0); i <= num; i++ {
		ret = append(ret, Point{X: p.X + i*sx, Y: p.Y + i*sy})
	}

	return ret
}

func (p *Point) Neighbour(b Point) bool {
	dx, dy := Abs(b.X-p.X), Abs(b.Y-p.Y)
	return dx <= 1 && dy <= 1
}

func IPow(base, exp int64) int64 {
	ret := int64(1)

	for {
		if exp&1 > 0 {
			ret *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return ret
}

func Atoi(a string) int {
	v, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalln(err)
	}

	return v
}

func Atoi64(a string) int64 {
	v, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return v
}

func Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for e := range seq {
				if !yield(e) {
					return
				}
			}
		}
	}
}
