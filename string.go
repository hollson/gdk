// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const BLANK = ""

// TrimSpace 修剪所有换行和空格，如：压缩整段格式化的SQL
func TrimSpace(str string) string {
	if len(str) == 0 {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return strings.TrimSpace(reg.ReplaceAllString(str, " "))
}

// Pascal 转换为帕斯卡命名
//  如: userName => UserName
//     user_name => UserName
func Pascal(title string) string {
	arr := strings.FieldsFunc(title, func(c rune) bool { return c == '_' })
	// RangeToTitle(arr)
	RangeStringsFunc(arr, func(s string) string { return strings.Title(s) })
	return strings.Join(arr, BLANK)
}

// RangeToTitle 遍历并将集合元素转换为Title格式
func RangeToTitle(arr []string) {
	for k, v := range arr {
		arr[k] = strings.Title(v)
	}
}

// RangeStringsFunc 遍历处理集合成员
func RangeStringsFunc(arr []string, f func(string) string) {
	for k, v := range arr {
		arr[k] = f(v)
	}
}

// HasAny raw中是否包含sub...中的任意子串
//  如: HasSub("datetime","date") => true
//     HasSub("datetime","time") => true
func HasAny(src string, sub ...string) bool {
	for _, v := range sub {
		if strings.Contains(src, v) {
			return true
		}
	}
	return false
}

// CharCount 统计字符个数
//  与javascript中的string.length一致
func CharCount(str string) int {
	return utf8.RuneCountInString(str)
}
