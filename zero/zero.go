// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package zero 类型缺省值
package zero

// True 获取零值指针
func True() *bool {
	var x = true
	return &x
}

// False 获取零值指针
func False() *bool {
	var x = false
	return &x
}

// Int 获取零值指针
func Int() *int {
	var x = 0
	return &x
}

// Int32 获取零值指针
func Int32() *int32 {
	var x int32 = 0
	return &x
}

// Int64 获取零值指针
func Int64() *int64 {
	var x int64 = 0
	return &x
}

// Float32 获取零值指针
func Float32() *float32 {
	var x float32 = 0
	return &x
}

// Float64 获取零值指针
func Float64() *float64 {
	var x float64 = 0
	return &x
}

// String 获取零值指针
func String() *string {
	var x string
	return &x
}

// Struct 获取零值指针
//  空结构体共享 runtime.zerobase
func Struct() *struct{} {
	return &struct{}{}
}

// IntArray 获取零值指针
func IntArray() *[]int {
	var x []int
	return &x
}

// Int32Array 获取零值指针
func Int32Array() *[]int32 {
	var x []int32
	return &x
}

// Int64Array 获取零值指针
func Int64Array() *[]int64 {
	var x []int64
	return &x
}

// Float32Array 获取零值指针
func Float32Array() *[]float32 {
	var x []float32
	return &x
}

// Float64Array 获取零值指针
func Float64Array() *[]float64 {
	var x []float64
	return &x
}

// StringArray 获取零值指针
func StringArray() *[]string {
	var x []string
	return &x
}
