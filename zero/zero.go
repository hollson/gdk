// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package zero

// 变量零值，用于快捷变量取址
var (
	True                = true
	False               = false
	Int                 = 0
	Int32       int32   = 0
	Int64       int64   = 0
	Float32     float32 = 0.0
	Float64             = 0.0
	String              = ""
	Struct              = struct{}{}
	ArrayInt            = [0]int{}
	ArrayInt32          = [0]int32{}
	ArrayInt64          = [0]int64{}
	ArrayString         = [0]string{}
)

func test() {
	var a *string = &String
	var b *int64 = &Int64
	println(a, b)
}
