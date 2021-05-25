// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package set

// 倒序翻转(byte), 如:[2,9,5,3] => [3,5,9,2]
func ReverseByte(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 倒序翻转(int),如:[2,9,5,3] => [3,5,9,2]
func ReverseInt(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 倒序翻转(int32), 如:[2,9,5,3] => [3,5,9,2]
func ReverseInt32(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 倒序翻转(int64), 如:[2,9,5,3] => [3,5,9,2]
func ReverseInt64(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 倒序翻转(float32), 如:[2,9,5,3] => [3,5,9,2]
func ReverseFloat32(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 倒序翻转(float64), 如:[2,9,5,3] => [3,5,9,2]
func ReverseFloat64(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 倒序翻转(string), 如:["d","a","f"] => ["f","a","d"]
func ReverseString(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
