package di

import (
	"go.uber.org/dig"
)

type DI struct {
	container *dig.Container
}

func New(options ...dig.Option) *DI {
	return &DI{
		container: dig.New(options...),
	}
}

func (d *DI) Provide(constructors ...Constructor) {
	for _, constructor := range constructors {
		if err := d.container.Provide(constructor.constructor, constructor.options...); err != nil {
			panic(err)
		}
	}
}

func (d *DI) Invoke(invoke InvokeFunction) error {
	return d.container.Invoke(invoke.function, invoke.options...)
}
