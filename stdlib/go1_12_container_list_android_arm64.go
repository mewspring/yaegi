// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports container/list'. DO NOT EDIT.

import (
	"container/list"
	"reflect"
)

func init() {
	Symbols["container/list"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(list.New),

		// type definitions
		"Element": reflect.ValueOf((*list.Element)(nil)),
		"List":    reflect.ValueOf((*list.List)(nil)),

		// interface wrapper definitions

	}
}
