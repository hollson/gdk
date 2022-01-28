// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"time"
)

// 三目运算(ternary)
func T3(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}

// 中心化等精简雪花算法
var last int64
// 根据时间戳生成长度8的字符串雪花编码
func ShortSnow() string {
	var sid, now int64 = 0, time.Now().Unix()
	if now > last {
		last, sid = now, now
	} else {
		last++
		sid = last
	}
	return fmt.Sprintf("%X", sid)
}


