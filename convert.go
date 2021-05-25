// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
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

// 将字符串转换为bool类型，参数错误时返回false
//  1,t,T,TRUE, true, True  => true
//  0,f,F,FALSE,false,False => false
func ToBool(b string) (ret bool) {
	ret, err := strconv.ParseBool(b)
	if err != nil {
		return false
	}
	return
}
