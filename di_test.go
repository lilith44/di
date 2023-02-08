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

type d struct {
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

func newD(b *b) *d {
	return &d{b: b}
}

func TestNew(t *testing.T) {
	di := New()

	di.Provide(
		NewConstructor(newA),
		NewConstructor(newB),
		NewConstructor(newC),
		NewConstructor(newD),
	)

	di.Prepare(
		NewInvokeFunction(newC),
		NewInvokeFunction(newD),
	)

	di.Invoke()
}
