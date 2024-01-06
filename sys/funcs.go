package sys

import (
	"errors"
	"reflect"
)

var (
	// ErrParamsNotAdapted  params length invalid
	ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

// Funcs bundle of functions
type Funcs map[string]reflect.Value

// NewFuncs function maps
func NewFuncs() Funcs {
	return make(Funcs, 0)
}

// Bind the function with the given function name
func (f Funcs) Bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()
	// 根据签名设置名称
	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	f[name] = v
	return
}

// Call the function with the given name and params
func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := f[name]; !ok {
		err = errors.New(name + " does not exist.")
		return
	}

	if len(params) != f[name].Type().NumIn() {
		err = ErrParamsNotAdapted
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in)
	return
}
