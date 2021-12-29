// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"unsafe"
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
	fmt.Println(ToBool("NULL"))
	fmt.Println(ToBool("nil"))
}

func TestB2i(t *testing.T) {
	fmt.Println(I2Bytes(math.MaxUint64))
	fmt.Println(I2Bytes(256))
	fmt.Println(I2Bytes(256 * 256)) // 相当于256进制

	fmt.Println(Bytes2I([8]byte{6: 4}))
	fmt.Println(Bytes2IAuto([]byte{6: 4}))
}

func slice2arrayWithHack() {
	var b = []int{11, 12, 13}
	var a = *(*[3]int)(unsafe.Pointer(&b[0]))
	a[1] += 10
	fmt.Printf("%v\n", b) // [11 12 13]
}

func TestByte8(t *testing.T) {
	ExampleByte8()
}

func ExampleByte8() {
	fmt.Println(Byte8([]byte{'a', 'b', 'c'}))
	fmt.Println(Byte8([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	// out:
	// [97 98 99 0 0 0 0 0]
	// [1 2 3 4 5 6 7 8]
}
