// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports crypto/md5'. DO NOT EDIT.

import (
	"crypto/md5"
	"reflect"
)

func init() {
	Value["crypto/md5"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize": reflect.ValueOf(md5.BlockSize),
		"New":       reflect.ValueOf(md5.New),
		"Size":      reflect.ValueOf(md5.Size),
		"Sum":       reflect.ValueOf(md5.Sum),

		// type definitions

	}
}