// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports os/signal'. DO NOT EDIT.

import (
	"os/signal"
	"reflect"
)

func init() {
	Symbols["os/signal"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Ignore":  reflect.ValueOf(signal.Ignore),
		"Ignored": reflect.ValueOf(signal.Ignored),
		"Notify":  reflect.ValueOf(signal.Notify),
		"Reset":   reflect.ValueOf(signal.Reset),
		"Stop":    reflect.ValueOf(signal.Stop),

		// type definitions

		// interface wrapper definitions

	}
}
