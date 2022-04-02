// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/hollson/gdk/color"
)

func main() {
	txt:=color.Text("你好呀",color.FontBlue,color.BackMagenta)
	fmt.Println(txt)
}
