package aoc

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"golang.org/x/exp/constraints"
)

// Union return a slice of the unique union of
// the input slices
func Union[T comparable](x []T, y []T) []T {
	hash := map[T]struct{}{}

	for _, v := range x {
		hash[v] = struct{}{}
	}
	for _, v := range y {
		hash[v] = struct{}{}
	}

	return slices.Collect(maps.Keys(hash))
}

func Intersection[T comparable](x ...[]T) []T {
	hash := make(map[T]struct{})
	output := make(map[T]struct{})

	for _, v := range x[0] {
		hash[v] = struct{}{}
	}

	for _, y := range x[1:] {
		for _, v := range y {
			if _, ok := hash[v]; ok {
				output[v] = struct{}{}
			}
		}
		hash = output
		output = make(map[T]struct{})
	}

	return slices.Collect(maps.Keys(hash))
}

// Select returns a slice of the input values that return
// true when evalueted by f
func Select[T any](in []T, f func(i T) bool) []T {
	out := make([]T, 0)
	for _, v := range in {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}

func SelectMap[T comparable, V any](in map[T]V, f func(i T, j V) bool) map[T]V {
	out := make(map[T]V)
	for k, v := range in {
		if f(k, v) {
			out[k] = v
		}
	}

	return out
}

// Select returns a slice of the input values that return
// false when evalueted by f
func Reject[T any](in []T, f func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range in {
		if !f(v) {
			out = append(out, v)
		}
	}

	return out
}

// Map takes an input slice and a mapping function and returns
// a slice of values as modified by f
func Map[T any, V any](in []T, f func(T) V) []V {
	out := make([]V, len(in))
	for i, v := range in {
		out[i] = f(v)
	}

	return out
}

func Each[T any](in []T, f func(T)) {
	for _, v := range in {
		f(v)
	}
}

func EachMap[T comparable, V any](in map[T]V, f func(T, V)) {
	for k, v := range in {
		f(k, v)
	}
}

func MapWithIndex[T any, V any](in []T, f func(int, T) V) []V {
	out := make([]V, len(in))
	for i, v := range in {
		out[i] = f(i, v)
	}

	return out
}

func MapM[T comparable, V any](in []T, f func(T) (T, V)) map[T]V {
	ret := make(map[T]V)
	for _, x := range in {
		k, v := f(x)
		ret[k] = v
	}

	return ret
}

func MapMap[T comparable, V any, X comparable, Y any](in map[T]V, f func(T, V) (X, Y)) map[X]Y {
	ret := make(map[X]Y)
	for k, v := range in {
		nk, nv := f(k, v)
		ret[nk] = nv
	}

	return ret
}

func Carve[T any](in []T, i, j int) []T {
	ret := in[:i]
	ret = append(ret, in[j:]...)

	return ret
}

// Reduce takes an input slice, and initial value and a reduce function
// it passes every value in the input slice to the reduceFn and returns
// the final value returned by reduceFn
func Reduce[T any, A any](in []T, initial A, reduceFn func(A, T) A) A {
	for _, v := range in {
		initial = reduceFn(initial, v)
	}

	return initial
}

// Cons take an input slice and an offset and loops though the slice
// making a tuple of index and index+offset
func Cons[T any](in []T, offset int) [][2]T {
	out := make([][2]T, 0)
	for i := 0; i < len(in)-offset; i++ {
		out = append(out, [2]T{in[i], in[i+offset]})
	}

	return out
}

// Zip take a variadic number of input slices and produces a single output
// slice of tuples of the ith value of each input slice
func Zip[T any](in ...[]T) [][]T {
	out := make([][]T, len(in[0]))
	for _, slice := range in {
		for i, elem := range slice {
			out[i] = append(out[i], elem)
		}
	}

	return out
}

// All returns a bool indicating if all values in the input slice
// produce a true value when evaluated by allFn
func All[T any](in []T, allFn func(T) bool) bool {
	out := true
	for _, v := range in {
		out = out && allFn(v)
	}

	return out
}

// Any returns a bool indicating if any value in the input slice
// produce a true value when evaluated by anyFn
func Any[T any](in []T, anyFn func(T) bool) bool {
	out := false
	for _, v := range in {
		out = out || anyFn(v)
	}

	return out
}

// Count returns an int of the number if values in the input
// slice that return true when evaluated by countFn
func Count[T any](in []T, countFn func(T) bool) int {
	out := 0
	for _, v := range in {
		if countFn(v) {
			out++
		}
	}

	return out
}

// Includes returns true if any value in the input slice
// matches value
func Includes[T comparable](in []T, value T) bool {
	for _, v := range in {
		if v == value {
			return true
		}
	}

	return false
}

// Max returns the maximum value in the input slice
func Max[T constraints.Ordered](in []T) T {
	if len(in) == 1 {
		return in[0]
	}

	max := in[0]
	for _, v := range in[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// MaxWithIndex returns the index of, and the maximum value in the input slice
func MaxWithIndex[T constraints.Ordered](in []T) (int, T) {
	if len(in) == 1 {
		return 0, in[0]
	}

	max := in[0]
	index := 0
	for i, v := range in[1:] {
		if v > max {
			max = v
			index = i + 1
		}
	}

	return index, max
}

// Min returns the minimum value in the input slice
func Min[T constraints.Ordered](in []T) T {
	if len(in) == 1 {
		return in[0]
	}

	min := in[0]
	for _, v := range in[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// MinWithIndex returns the minimum value in the input slice
func MinWithIndex[T constraints.Ordered](in []T) (int, T) {
	if len(in) == 1 {
		return 0, in[0]
	}

	min := in[0]
	index := 0
	for i, v := range in[1:] {
		if v < min {
			min = v
			index = i + 1
		}
	}

	return index, min
}

// MinMax returns two values, the minimum and maximum value
// in the input slice
func MinMax[T constraints.Ordered](in []T) (T, T) {
	if len(in) == 1 {
		return in[0], in[0]
	}
	min, max := in[0], in[0]
	for _, v := range in[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// Partition takes an input slice and returns two slices using
// partFn to decide which it goes in, the first return value is
// the slice of values that return true when evaluated by partFn
// the second return is those that return false
func Partition[T any](in []T, partFn func(T) bool) ([]T, []T) {
	outT := make([]T, 0)
	outF := make([]T, 0)

	for _, v := range in {
		if partFn(v) {
			outT = append(outT, v)
		} else {
			outF = append(outF, v)
		}
	}

	return outT, outF
}

// Sum returns the sum of values in the input slice
func Sum[T constraints.Integer | constraints.Float](in []T) T {
	var out T
	for _, v := range in {
		out += v
	}
	return out
}

// Product returns the product of values in the input slice
func Product[T constraints.Integer | constraints.Float](in []T) T {
	var out T = 1

	for _, v := range in {
		out *= v
	}

	return out
}

// Combinations return a slice of tuples of all 2 value combinations
// of the input slice
func Combinations[T any](in []T) [][2]T {
	out := make([][2]T, 0)
	for i, x := range in {
		for _, y := range in[i+1:] {
			out = append(out, [2]T{x, y})
		}
	}

	return out
}

func Permutations[T any](in []T) [][]T {
	var helper func([]T, int)
	var ret [][]T

	helper = func(in []T, n int) {
		if n == 1 {
			ret = append(ret, slices.Clone(in))
		} else {
			for i := 0; i < n; i++ {
				helper(in, n-1)
				if n%2 == 1 {
					tmp := in[i]
					in[i] = in[n-1]
					in[n-1] = tmp
				} else {
					tmp := in[0]
					in[0] = in[n-1]
					in[n-1] = tmp
				}
			}
		}
	}
	helper(in, len(in))
	return ret
}

func Join[T any](in []T, sep string) string {
	var b strings.Builder
	for _, v := range in {
		fmt.Fprintf(&b, "%s%v", sep, v)
	}

	return b.String()[1:]
}

func Sort[T constraints.Ordered](in []T) []T {
	slices.Sort(in)

	return slices.Clone(in)
}

func Reverse[T any](in []T) []T {
	ret := make([]T, len(in))

	for i := 0; i < len(in); i++ {
		ret[len(in)-1-i] = in[i]
	}

	return ret
}

// Tally takes a slice of comparable items and returns a map with a tally
// of the number of times each item appears in the slice
func Tally[T comparable](in []T) map[T]int {
	ret := make(map[T]int)
	for _, v := range in {
		ret[v] += 1
	}

	return ret
}

// KeyWithMaxValue takes a map and returns the key in the map that is associated
// with the maximum value. As maps are read in random order if the maximum value
// is shared the returned k is not deterministic
func KeyWithMaxValue[T comparable, V constraints.Ordered](in map[T]V) T {
	if len(in) == 1 {
		return slices.Collect(maps.Keys(in))[0]
	}

	var maxK T
	var maxV V
	for k, v := range in {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}

	return maxK
}

func KeyWithValue[T comparable, V comparable](in map[T]V, val V) T {
	var e T
	for k, v := range in {
		if v == val {
			return k
		}
	}

	return e
}

func Chunk[T any](in []T, size int) [][]T {
	var ret [][]T
	for i := 0; i < len(in)/size; i++ {
		skip := size
		rem := len(in) - i*size
		if rem < skip {
			skip = rem
		}
		ret = append(ret, in[i*size:i*size+skip])
	}

	return ret
}

func SubSlice[T comparable](a, b []T) []T {
	var ret []T
	for _, v := range a {
		if !slices.Contains(b, v) {
			ret = append(ret, v)
		}
	}

	return ret
}
