// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

// 三目运算
func T3(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}