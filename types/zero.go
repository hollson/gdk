// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

// True default value pointer
func True() *bool {
	var x = true
	return &x
}

// False default value pointer
func False() *bool {
	var x = false
	return &x
}

// Int default value pointer
func Int() *int {
	var x = 0
	return &x
}

// Int32 default value pointer
func Int32() *int32 {
	var x int32 = 0
	return &x
}

// Int64 default value pointer
func Int64() *int64 {
	var x int64 = 0
	return &x
}

// Float32 default value pointer
func Float32() *float32 {
	var x float32 = 0
	return &x
}

// Float64 default value pointer
func Float64() *float64 {
	var x float64 = 0
	return &x
}

// String default value pointer
func String() *string {
	var x string
	return &x
}

// Struct default value pointer
//
//	空结构体共享 runtime.zerobase
func Struct() *struct{} {
	return &struct{}{}
}

// IntArray default value pointer
func IntArray() *[]int {
	var x []int
	return &x
}

// Int32Array default value pointer
func Int32Array() *[]int32 {
	var x []int32
	return &x
}

// Int64Array default value pointer
func Int64Array() *[]int64 {
	var x []int64
	return &x
}

// Float32Array default value pointer
func Float32Array() *[]float32 {
	var x []float32
	return &x
}

// Float64Array default value pointer
func Float64Array() *[]float64 {
	var x []float64
	return &x
}

// StringArray default value pointer
func StringArray() *[]string {
	var x []string
	return &x
}
