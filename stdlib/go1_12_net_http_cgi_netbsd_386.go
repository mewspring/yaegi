// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports net/http/cgi'. DO NOT EDIT.

import (
	"net/http/cgi"
	"reflect"
)

func init() {
	Symbols["net/http/cgi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Request":        reflect.ValueOf(cgi.Request),
		"RequestFromMap": reflect.ValueOf(cgi.RequestFromMap),
		"Serve":          reflect.ValueOf(cgi.Serve),

		// type definitions
		"Handler": reflect.ValueOf((*cgi.Handler)(nil)),

		// interface wrapper definitions

	}
}
