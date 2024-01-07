// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sys

// T3 三目运算(ternary)
func T3(b bool, x, y any) any {
	if b {
		return x
	}
	return y
}
