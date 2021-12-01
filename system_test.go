// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"testing"
)

func TestT3(t *testing.T) {
	a, b := 2, 3
	max := T3(a > b, a, b).(int)
	println(max)
}

