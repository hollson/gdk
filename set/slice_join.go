// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package set

import (
	"strconv"
	"strings"
)

func JoinInt32(arr []int32, splitter ...string) string {
	var b strings.Builder
	length := len(arr)
	b.Grow(length) // 预分配容量
	for i := 0; i < length; i++ {
		b.WriteString(strconv.Itoa(int(arr[i])))
		b.WriteString(splitter[0])
	}
	return b.String()
}

func JoinInt64(arr []int64, splitter ...string) string {
	var b strings.Builder
	length := len(arr)
	b.Grow(length) // 预分配容量
	for i := 0; i < length; i++ {
		b.WriteString(strconv.Itoa(int(arr[i])))
		b.WriteString(splitter[0])
	}
	return b.String()
}


func JoinInt(arr []int32, splitter ...string) string {
	var b strings.Builder
	length := len(arr)
	b.Grow(length) // 预分配容量
	for i := 0; i < length; i++ {
		b.WriteString(strconv.Itoa(int(arr[i])))
		b.WriteString(splitter[0])
	}
	return b.String()
}
