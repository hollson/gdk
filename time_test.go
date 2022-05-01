// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	fmt.Println(ParseDate("2022-12-12"))
	fmt.Println(ParseTime("2022-12-12 22:22:22"))
}

func TestTime1970(t *testing.T) {
	fmt.Println(Time1970())
}

func Test1970(t *testing.T) {
	fmt.Println(time.UnixMicro(0))
}
