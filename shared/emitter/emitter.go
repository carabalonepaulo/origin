package emitter

import "github.com/carabalonepaulo/origin/shared/slab"

const minSliceSize = 0

type Emitter struct {
	evs []slab.Slab[func(any)]
}

func Init(evs int) Emitter {
	s := Emitter{}
	s.evs = make([]slab.Slab[func(any)], evs)

	for i := 0; i < evs; i++ {
		s.evs[i] = slab.Init[func(any)](minSliceSize)
	}

	return s
}

func (e *Emitter) Emit(ev int, arg any) {
	iter := e.evs[ev].Iter()
	for iter.Next() {
		(*iter.Value())(arg)
	}
}

func (e *Emitter) On(ev int, cb func(any)) func() {
	k := e.evs[ev].Insert(cb)
	return func() { e.evs[ev].Remove(k) }
}

func (e *Emitter) Once(ev int, cb func(any)) {
	ve := e.evs[ev].VacantEntry()
	k := ve.Key()
	ve.Insert(func(arg any) {
		cb(arg)
		e.evs[ev].Remove(k)
	})
}
