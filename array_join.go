// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// 更多详情，请参考：https://github.com/hollson/gdk

package gdk

import (
	"fmt"
	"strconv"
	"strings"
)

// JoinByte 拼接切片
//  splitter：拼接符（默认为逗号）
func JoinByte(arr []byte, splitter ...string) string {
	if len(arr) == 0 {
		return ""
	}
	if len(splitter) == 0 {
		splitter = []string{","}
	}
	var (
		length = len(arr)
		b      strings.Builder
	)
	b.Grow(length) // 预分配容量

	for k, v := range arr {
		b.WriteByte(v)
		if k < length-1 {
			b.WriteString(splitter[0])
		}
	}
	return b.String()
}

// JoinInt 将切片拼接成以splitter(默认逗号)分割等字符串
func JoinInt(arr []int, splitter ...string) string {
	if len(arr) == 0 {
		return ""
	}
	if len(splitter) == 0 {
		splitter = []string{","}
	}

	var (
		length = len(arr)
		b      strings.Builder
	)
	b.Grow(length) // 预分配容量

	for k, v := range arr {
		b.WriteString(strconv.Itoa(v))
		if k < length-1 {
			b.WriteString(splitter[0])
		}
	}
	return b.String()
}

// 将切片拼接成以splitter(默认逗号)分割等字符串
func JoinInt32(arr []int32, splitter ...string) string {
	if len(arr) == 0 {
		return ""
	}
	if len(splitter) == 0 {
		splitter = []string{","}
	}

	var (
		length = len(arr)
		b      strings.Builder
	)
	b.Grow(length) // 预分配容量

	for k, v := range arr {
		b.WriteString(strconv.Itoa(int(v)))
		if k < length-1 {
			b.WriteString(splitter[0])
		}
	}
	return b.String()
}

func JoinInt64(arr []int64, splitter ...string) string {
	if len(arr) == 0 {
		return ""
	}
	if len(splitter) == 0 {
		splitter = []string{","}
	}

	var (
		length = len(arr)
		b      strings.Builder
	)
	b.Grow(length) // 预分配容量

	for k, v := range arr {
		b.WriteString(strconv.Itoa(int(v)))
		if k < length-1 {
			b.WriteString(splitter[0])
		}
	}
	return b.String()
}

// 将切片拼接成「sql in」语句
//  如：{1,3,5,7,9} ==> "1,3,5,7,9"
func SQLInInt64(items ...int64) (inSql string) {
	if len(items) == 0 {
		return "''"
	}
	for _, v := range items {
		inSql = fmt.Sprintf("%s,%d", inSql, v)
	}
	return strings.TrimLeft(inSql, ",")
}

// 将切片拼接成「sql in」语句
//  如：{"aa","bb","cc"} ==> "'1','3','5','7','9'"
func SQLInString(items ...string) (inSql string) {
	if len(items) == 0 {
		return "''"
	}
	for _, v := range items {
		inSql = fmt.Sprintf("%s,'%s'", inSql, v)
	}
	return strings.TrimLeft(inSql, ",")
}

// 将切片拼接成「sql in」语句
//  如：{1,3,5,7,9} ==> "1,3,5,7,9"
func SQLInInt(items ...int) (inSql string) {
	if len(items) == 0 {
		return "''"
	}
	for _, v := range items {
		inSql = fmt.Sprintf("%s,%d", inSql, v)
	}
	return strings.TrimLeft(inSql, ",")
}

// 将切片拼接成「sql in」语句
//  如：{1,3,5,7,9} ==> "1,3,5,7,9"
func SQLInInt32(items ...int32) (inSql string) {
	if len(items) == 0 {
		return "''"
	}
	for _, v := range items {
		inSql = fmt.Sprintf("%s,%d", inSql, v)
	}

	return strings.TrimLeft(inSql, ",")
}

// func JoinFunc(items []interface{}) {
//
// 	strings.IndexFunc()
// }
