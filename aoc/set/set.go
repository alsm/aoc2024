package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	s := make(map[T]struct{})

	return s
}

func NewFromSlice[T comparable](x []T) Set[T] {
	s := make(map[T]struct{})
	for _, v := range x {
		s[v] = struct{}{}
	}

	return s
}

func (s Set[T]) Contains(x T) bool {
	_, ok := s[x]

	return ok
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Delete(x T) {
	delete(s, x)
}

func (s Set[T]) Clear() {
	s = make(map[T]struct{})
}

func (s Set[T]) Sub(o Set[T]) {
	for k := range o {
		delete(s, k)
	}
}

func (s Set[T]) SubA(o []T) {
	for _, k := range o {
		delete(s, k)
	}
}
