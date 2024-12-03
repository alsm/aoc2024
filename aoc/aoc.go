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

func Filter[T any](f func(T) bool, s iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func Filter2[T, U any](f func(T, U) bool, s iter.Seq2[T, U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for k, v := range s {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

func Map[In, Out any](f func(In) Out, s iter.Seq[In]) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for in := range s {
			if !yield(f(in)) {
				return
			}
		}
	}
}

func Map2[KIn, VIn, KOut, VOut any](f func(KIn, VIn) (KOut, VOut), s iter.Seq2[KIn, VIn]) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for k, v := range s {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func Reduce[Sum, V any](f func(Sum, V) Sum, sum Sum, s iter.Seq[V]) Sum {
	for v := range s {
		sum = f(sum, v)
	}
	return sum
}

func Reduce2[Sum, K, V any](f func(Sum, K, V) Sum, sum Sum, s iter.Seq2[K, V]) Sum {
	for k, v := range s {
		sum = f(sum, k, v)
	}
	return sum
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

type Zipped[V1, V2 any] struct {
	V1  V1
	Ok1 bool // whether V1 is present (if not, it will be zero)
	V2  V2
	Ok2 bool // whether V2 is present (if not, it will be zero)
}

func Zip[V1, V2 any](x iter.Seq[V1], y iter.Seq[V2]) iter.Seq[Zipped[V1, V2]] {
	return func(yield func(z Zipped[V1, V2]) bool) {
		next, stop := iter.Pull(y)
		defer stop()
		v2, ok2 := next()
		for v1 := range x {
			if !yield(Zipped[V1, V2]{v1, true, v2, ok2}) {
				return
			}
			v2, ok2 = next()
		}
		var zv1 V1
		for ok2 {
			if !yield(Zipped[V1, V2]{zv1, false, v2, ok2}) {
				return
			}
			v2, ok2 = next()
		}
	}
}

func Tally[T comparable](in []T) map[T]int {
	ret := make(map[T]int)
	for _, v := range in {
		ret[v] += 1
	}

	return ret
}

func Count[T any](f func(T) bool, in iter.Seq[T]) int {
	var ret int
	for v := range in {
		if f(v) {
			ret++
		}
	}

	return ret
}

func All[T any](f func(T) bool, in iter.Seq[T]) bool {
	ret := true
	for v := range in {
		ret = ret && f(v)
	}

	return ret
}

func Any[T any](f func(T) bool, in iter.Seq[T]) bool {
	ret := false
	for v := range in {
		ret = ret || f(v)
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
