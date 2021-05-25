// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"fmt"
	"strconv"
	"testing"
)

func TestToInt64(t *testing.T) {
	// 将16进制的字符串FF转换为10进制
	// base：进位制（2 进制到 36 进制）
	// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
	// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	fmt.Println(strconv.ParseInt("011", 16, 8))
}

func TestToBool(t *testing.T) {
	fmt.Println(ToBool("t"))
	fmt.Println(ToBool("F"))
	fmt.Println(ToBool("1"))
	fmt.Println(ToBool("true"))
	fmt.Println(ToBool("True"))
	fmt.Println(ToBool("F"))
	fmt.Println(ToBool("TRUE"))
}
