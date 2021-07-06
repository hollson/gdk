// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

// 倒序翻转(byte), 如:[9,5,2,7] => [7,2,5,9]
func ReverseBytes(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 倒序翻转(int),如:[9,5,2,7] => [7,2,5,9]
func ReverseInts(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 倒序翻转(int32), 如:[9,5,2,7] => [7,2,5,9]
func ReverseInt32s(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 倒序翻转(int64), 如:[9,5,2,7] => [7,2,5,9]
func ReverseInt64s(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 倒序翻转(float32), 如:[9,5,2,7] => [7,2,5,9]
func ReverseFloat32s(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 倒序翻转(float64), 如:[9,5,2,7] => [7,2,5,9]
func ReverseFloat64s(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 倒序翻转(string), 如:["a","r","e"] => ["e","r","a"]
func ReverseStrings(arr ...int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
