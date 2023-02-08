package di

import (
	"testing"
)

type a struct {
}

type b struct {
	a *a
}

type c struct {
	b *b
}

func newA() *a {
	return new(a)
}

func newB(a *a) *b {
	return &b{a: a}
}

func newC(b *b) *c {
	return &c{b: b}
}

func TestNew(t *testing.T) {
	di := New()

	di.Provide(
		NewConstructor(newA),
		NewConstructor(newB),
	)

	di.Invoke(NewInvokeFunction(newC))
}
