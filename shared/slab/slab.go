package slab

import (
	"github.com/carabalonepaulo/origin/shared/stack"
)

type Key int

type Slab[T any] struct {
	items     []T
	available stack.Stack[Key]
	used      []int
}

func Init[T any](c int) Slab[T] {
	st := stack.Init[Key](c)
	for i := 0; i < c; i++ {
		st.Push(Key(i))
	}

	return Slab[T]{
		items:     make([]T, c),
		available: st,
		used:      make([]int, 0, c),
	}
}

func (s *Slab[T]) Insert(v T) Key {
	if s.available.Empty() {
		s.grow()
	}
	k, _ := s.available.Pop()
	s.items[int(k)] = v
	s.used = append(s.used, int(k))
	return k
}

func (s *Slab[T]) Remove(k Key) (v T, ok bool) {
	s.available.Push(k)
	s.removeFromUsed(k)
	v = s.items[k]
	ok = true
	return
}

func (s *Slab[T]) Length() int {
	return len(s.used)
}

func (s *Slab[T]) Ref(k Key) *T {
	if s.findUsedKey(k) == -1 {
		return nil
	}
	return &s.items[k]
}

type Iter[T any] struct {
	src *Slab[T]
	i   int
}

func (i *Iter[T]) Value() *T {
	return &i.src.items[i.src.used[i.i]]
}

func (i *Iter[T]) Next() bool {
	i.i += 1
	return i.i < len(i.src.used)
}

func (s *Slab[T]) Iter() Iter[T] {
	return Iter[T]{src: s, i: -1}
}

type Entry[T any] struct {
	src *Slab[T]
	k   Key
	u   bool
}

func (s *Slab[T]) VacantEntry() Entry[T] {
	if s.available.Empty() {
		s.grow()
	}
	k, _ := s.available.Pop()

	return Entry[T]{
		src: s,
		k:   k,
	}
}

func (e *Entry[T]) Key() Key {
	return e.k
}

func (e *Entry[T]) Insert(v T) {
	if e.u {
		return
	}

	e.src.items[e.k] = v
	e.src.used = append(e.src.used, int(e.k))
	e.u = true
}

func (e *Entry[T]) Discard() {
	if e.u {
		return
	}

	e.src.removeFromUsed(e.k)
	e.src.available.Push(e.k)
	e.u = true
}

func (s *Slab[T]) grow() {
	oldCap := cap(s.items)
	if oldCap == 0 {
		oldCap = 1
	}

	t := make([]T, oldCap*2)
	copy(t, s.items)
	s.items = t

	for i := oldCap; i < cap(t); i++ {
		s.available.Push(Key(i))
	}
}

func (s *Slab[T]) findUsedKey(k Key) int {
	j := int(k)
	idx := -1
	for i := 0; i < len(s.used); i++ {
		if s.used[i] == j {
			idx = i
			break
		}
	}

	if idx == -1 {
		// log.Fatalf("Invalid key {%d}!\n", k)
		panic("")
	}
	return idx
}

func (s *Slab[T]) removeFromUsed(k Key) {
	idx := s.findUsedKey(k)
	t := s.used[idx]
	l := len(s.used) - 1
	s.used[idx] = s.used[l]
	s.used[l] = t
	s.used = s.used[0:l]
}
