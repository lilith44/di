package di

import "go.uber.org/dig"

type Constructor struct {
	constructor any
	options     []dig.ProvideOption
}

func NewConstructor(constructor any, options ...dig.ProvideOption) Constructor {
	return Constructor{
		constructor: constructor,
		options:     options,
	}
}
