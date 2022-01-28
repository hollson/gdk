// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

const BLANK = ""

// 修剪所有换行和空格，如：压缩整段格式化的SQL
func TrimSpace(str string) string {
	if len(str) == 0 {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return strings.TrimSpace(reg.ReplaceAllString(str, " "))
}

// 转换为帕斯卡命名
//  如: userName => UserName
//     user_name => UserName
func Pascal(title string) string {
	arr := strings.FieldsFunc(title, func(c rune) bool { return c == '_' })
	// RangeToTitle(arr)
	RangeStringsFunc(arr, func(s string) string { return strings.Title(s) })
	return strings.Join(arr, BLANK)
}

// 遍历并将集合元素转换为Title格式
func RangeToTitle(arr []string) {
	for k, v := range arr {
		arr[k] = strings.Title(v)
	}
}

// 遍历处理集合成员
func RangeStringsFunc(arr []string, f func(string) string) {
	for k, v := range arr {
		arr[k] = f(v)
	}
}

// raw中是否包含sub...中的任意子串
//  如: HasSub("datetime","date") => true
//     HasSub("datetime","time") => true
func HasAny(raw string, sub ...string) bool {
	for _, v := range sub {
		if strings.Contains(raw, v) {
			return true
		}
	}
	return false
}

// 统计字符个数(汉字:1，字母:0.5)
func GetChars(str string) int64 {
	var count float64 = 0
	for _, val := range str {
		if unicode.Is(unicode.Scripts["Han"], val) {
			count += 1
		} else {
			count += 0.5
		}
	}
	return int64(math.Ceil(count))
}

func GetStrLength(str string) float64 {
	var total float64
	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 1
		} else {
			total = total + 0.5
		}
	}
	return math.Ceil(total)
}
