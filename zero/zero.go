// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package zero

var (
	True                = true        // (全局)变量零值，用于快捷取指，请勿改变变量值
	False               = false       // (全局)变量零值，用于快捷取指，请勿改变变量值
	Int                 = 0           // (全局)变量零值，用于快捷取指，请勿改变变量值
	Int32       int32   = 0           // (全局)变量零值，用于快捷取指，请勿改变变量值
	Int64       int64   = 0           // (全局)变量零值，用于快捷取指，请勿改变变量值
	Float32     float32 = 0.0         // (全局)变量零值，用于快捷取指，请勿改变变量值
	Float64             = 0.0         // (全局)变量零值，用于快捷取指，请勿改变变量值
	String              = ""          // (全局)变量零值，用于快捷取指，请勿改变变量值
	Struct              = struct{}{}  // (全局)变量零值，用于快捷取指，请勿改变变量值
	ArrayInt            = [0]int{}    // (全局)变量零值，用于快捷取指，请勿改变变量值
	ArrayInt32          = [0]int32{}  // (全局)变量零值，用于快捷取指，请勿改变变量值
	ArrayInt64          = [0]int64{}  // (全局)变量零值，用于快捷取指，请勿改变变量值
	ArrayString         = [0]string{} // (全局)变量零值，用于快捷取指，请勿改变变量值
)

// usage
// func test() {
// 	var a *string = &String
// 	var b *int64 = &Int64
// 	println(a, b)
// }
