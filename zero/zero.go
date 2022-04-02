// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package zero 类型缺省值
package zero

// 显式声明
var (
	_true         bool     = true
	_false        bool     = false
	_int          int      = 0
	_int32        int32    = 0
	_int64        int64    = 0
	_float32      float32  = 0.0
	_float64      float64  = 0.0
	_string       string   = ""
	_struct       struct{} = struct{}{}
	_intArray     []int
	_int32Array   []int32
	_int64Array   []int64
	_float32Array []float32
	_float64Array []float64
	_stringArray  []string
)

// True 获取当前类型的零值指针
func True() *bool {
	return &_true
}

// False 获取当前类型的零值指针
func False() *bool {
	return &_false
}

// Int 获取当前类型的零值指针
func Int() *int {
	return &_int
}

// Int32 获取当前类型的零值指针
func Int32() *int32 {
	return &_int32
}

// Int64 获取当前类型的零值指针
func Int64() *int64 {
	return &_int64
}

// Float32 获取当前类型的零值指针
func Float32() *float32 {
	return &_float32
}

// Float64 获取当前类型的零值指针
func Float64() *float64 {
	return &_float64
}

// String 获取当前类型的零值指针
func String() *string {
	return &_string
}

// Struct 获取当前类型的零值指针
func Struct() *struct{} {
	return &_struct
}

// IntArray 获取当前类型的零值指针
func IntArray() *[]int {
	return &_intArray
}

// Int32Array 获取当前类型的零值指针
func Int32Array() *[]int32 {
	return &_int32Array
}

// Int64Array 获取当前类型的零值指针
func Int64Array() *[]int64 {
	return &_int64Array
}

// Float32Array 获取当前类型的零值指针
func Float32Array() *[]float32 {
	return &_float32Array
}

// Float64Array 获取当前类型的零值指针
func Float64Array() *[]float64 {
	return &_float64Array
}

// StringArray 获取当前类型的零值指针
func StringArray() *[]string {
	return &_stringArray
}
