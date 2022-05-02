// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"encoding/binary"
	"strconv"
)

// 将字符串转换为int64，转换失败时返回0
func ToInt(n string) (ret int) {
	ret, err := strconv.Atoi(n)
	
	if err != nil {
		return 0
	}
	return
}

// 将字符串转换为int64，转换失败时返回0
func ToInt64(num string) int64 {
	// 将16进制的字符串FF转换为10进制
	// base:进位制(2-36), 如果为0,则根据前缀判断(如0=>8;0x=>16)
	// bitSize：整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	if nt64, err := strconv.ParseInt(num, 10, 0); err != nil {
		return 0
	} else {
		return nt64
	}
}

// ToBool 将字符串转换为bool类型，参数错误时返回false
//  1,t,T,TRUE, true, True  => true
//  0,f,F,FALSE,false,False => false
func ToBool(b string) bool {
	ret, err := strconv.ParseBool(b)
	if err != nil {
		return false
	}
	return ret
}

// I2Bytes int转换为byte数组
func I2Bytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// Bytes2I byte数组转换为int
func Bytes2I(b [8]byte) uint64 {
	return binary.BigEndian.Uint64(b[:])
}

// Bytes2IAuto byte数组转换为int
func Bytes2IAuto(b []byte) uint64 {
	return binary.BigEndian.Uint64(Byte8(b))
}

// Byte8 将给定的切片转换为[8]byte类型
//  说明： 使用copy(to,from)函数自由填充目标切片或数组，如：
//         n:=copy(arr,slice) 或
//         n:=copy(slice,arr)
func Byte8(s []byte) []byte {
	var ret [8]byte
	copy(ret[:], s)
	return ret[:]
}
