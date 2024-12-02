package stack

type Stack[T any] struct {
	i int
	s []T
}

func Init[T any](minCap int) Stack[T] {
	return Stack[T]{
		i: 0,
		s: make([]T, minCap),
	}
}

func (s *Stack[T]) Empty() bool { return s.i == 0 }

func (s *Stack[T]) Push(v T) {
	s.ensureCap()

	s.s[s.i] = v
	s.i += 1
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if s.i <= 0 {
		ok = false
		return
	}

	s.i -= 1
	v = s.s[s.i]
	ok = true
	return
}

func (s *Stack[T]) Clear() {
	s.i = 0
}

type Iter[T any] struct {
	src *Stack[T]
	i   int
}

func (i *Iter[T]) Value() (v T) {
	v = i.src.s[i.i]
	return
}

func (i *Iter[T]) Next() bool {
	i.i += 1
	return i.i < i.src.i
}

func (s *Stack[T]) Iter() Iter[T] {
	return Iter[T]{src: s, i: -1}
}

func (s *Stack[T]) ensureCap() {
	oc := cap(s.s)
	if oc == 0 {
		oc = 1
	}

	if s.i+1 < oc {
		return
	}

	t := make([]T, int(oc*2))
	copy(t, s.s)
	s.s = t
}
