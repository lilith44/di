package di

import "go.uber.org/dig"

type InvokeFunction struct {
	function any
	options  []dig.InvokeOption
}

func NewInvokeFunction(function any, options ...dig.InvokeOption) InvokeFunction {
	return InvokeFunction{
		function: function,
		options:  options,
	}
}
