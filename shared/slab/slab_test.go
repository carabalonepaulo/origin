package slab

import "testing"

func TestSlab(t *testing.T) {
	s := Init[int](3)
	s.Insert(10)
	b := s.Insert(20)
	s.Insert(30)

	if v, ok := s.Remove(b); !ok || v != 20 {
		t.Fail()
	}

	iter := s.Iter()

	if !iter.Next() {
		t.Fail()
	}
	if v := iter.Value(); *v != 10 {
		t.Fail()
	}

	if !iter.Next() {
		t.Fail()
	}
	if v := iter.Value(); *v != 30 {
		t.Fail()
	}

	if iter.Next() {
		t.Fail()
	}
}

func TestEntry(t *testing.T) {
	s := Init[int](3)
	e := s.VacantEntry()

	if e.k != 2 {
		t.Fail()
	}
	e.Insert(10)
	if s.Length() != 1 {
		t.Fail()
	}

	// can't use twice
	e.Insert(10)
	if s.Length() != 1 {
		t.Fail()
	}

	iter := s.Iter()
	if !iter.Next() {
		t.Fail()
	}
	if v := iter.Value(); *v != 10 {
		t.Fail()
	}
}

func TestSlabWithZeroCap(t *testing.T) {
	s := Init[int](0)
	s.Insert(10)
}

func TestPointerToValue(t *testing.T) {
	type Data struct{ a int }
	s := Init[Data](1)
	k := s.Insert(Data{a: 10})

	r := s.Ref(k)
	r.a = 20

	i := s.Iter()
	i.Next()
	if i.Value().a != 20 {
		t.Fail()
	}
}
