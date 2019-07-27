// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports net/textproto'. DO NOT EDIT.

import (
	"net/textproto"
	"reflect"
)

func init() {
	Symbols["net/textproto"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CanonicalMIMEHeaderKey": reflect.ValueOf(textproto.CanonicalMIMEHeaderKey),
		"Dial":                   reflect.ValueOf(textproto.Dial),
		"NewConn":                reflect.ValueOf(textproto.NewConn),
		"NewReader":              reflect.ValueOf(textproto.NewReader),
		"NewWriter":              reflect.ValueOf(textproto.NewWriter),
		"TrimBytes":              reflect.ValueOf(textproto.TrimBytes),
		"TrimString":             reflect.ValueOf(textproto.TrimString),

		// type definitions
		"Conn":          reflect.ValueOf((*textproto.Conn)(nil)),
		"Error":         reflect.ValueOf((*textproto.Error)(nil)),
		"MIMEHeader":    reflect.ValueOf((*textproto.MIMEHeader)(nil)),
		"Pipeline":      reflect.ValueOf((*textproto.Pipeline)(nil)),
		"ProtocolError": reflect.ValueOf((*textproto.ProtocolError)(nil)),
		"Reader":        reflect.ValueOf((*textproto.Reader)(nil)),
		"Writer":        reflect.ValueOf((*textproto.Writer)(nil)),

		// interface wrapper definitions

	}
}
