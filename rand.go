// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	NumberSymbol = "0123456789"
	LetterSymbol = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	PassSymbol   = "!@#$%^&*()[]{}+-*/_=."
)

var _rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 随机生成字符串
//  n: 生成字符串个数
//  letter: 基本字符组,即生成结果等必要数据源
func randStr(n int, letter string) string {
	s := []byte(letter)
	var buf bytes.Buffer
	buf.Grow(n)
	for i := 0; i < n; i++ {
		buf.WriteByte(s[_rand.Intn(strings.Count(letter, "")-1)])
	}
	return buf.String()
}

// RandNumStr 随机数字串,n为生成长度
func RandNumStr(n int) string {
	return randStr(n, NumberSymbol)
}

// RandString 随机字符串(包含数字和大小写字母),n为生成长度
func RandString(n int) string {
	return randStr(n, fmt.Sprintf("%s%s", NumberSymbol, LetterSymbol))
}

// RandPassword 随机密码串(包含数字、大小写字母和特殊符号),n为生成长度
func RandPassword(n int) string {
	return randStr(n, fmt.Sprintf("%s%s%s", NumberSymbol, LetterSymbol, PassSymbol))
}

// RandSequence 带有时间序列的随机数字串，即14位时间字符+N个随机字符，
// 如：20060102150405+99999
func RandSequence(n int) string {
	return time.Now().Format("20060102150405") + RandNumStr(n)
}

// RandScopeInt 生成m到n以内的随机数，左闭右开区间[m,n)
func RandScopeInt(m, n int) int {
	return _rand.Intn(n-m) + m
}

// RandInt 生成n以内的随机数
func RandInt(n int) int {
	return _rand.Intn(n)
}

// RandInt64 生成n以内的随机数
func RandInt64(n int64) int64 {
	return _rand.Int63n(n)
}

// Deprecated: 推荐使用 RandInt
func RandIntX(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
