// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package snow

import (
	"fmt"
	"time"
)

// 中心化的精简雪花算法
var last int64

// Create 根据时间戳生成长度8的字符串雪花编码
func Create() string {
	var sid, now int64 = 0, time.Now().Unix()
	if now > last {
		last, sid = now, now
	} else {
		last++ // fixme 原子操作
		sid = last
	}

	return fmt.Sprintf("%X", sid)
}
