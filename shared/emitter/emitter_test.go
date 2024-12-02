package emitter

import "testing"

func TestEmitter(t *testing.T) {
	a, b := false, false
	e := Init(2)
	da := e.On(0, func(_ any) { a = true })
	db := e.On(1, func(_ any) { b = true })
	e.Emit(0, nil)
	if !a {
		t.Fail()
	}
	e.Emit(1, nil)
	if !b {
		t.Fail()
	}
	a, b = false, false
	da()
	db()
	e.Once(0, func(_ any) { a = true })
	e.Once(1, func(_ any) { b = true })
	e.Emit(0, nil)
	if !a {
		t.Fail()
	}
	e.Emit(1, nil)
	if !b {
		t.Fail()
	}
	a, b = false, false
	e.Emit(0, nil)
	if a {
		t.Fail()
	}
	e.Emit(1, nil)
	if b {
		t.Fail()
	}
}
