// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sys_test

import (
	"fmt"
	"testing"

	"github.com/hollson/gdk"
	"github.com/hollson/gdk/sys"
)

func TestName(t *testing.T) {
	fmt.Println(rune('👩'))
	fmt.Println(011)
	fmt.Println(string(rune(0x274C)))
}

func TestCallerMust(t *testing.T) {
	ExampleCallerMust()
}

func ExampleCallerMust() {
	c := sys.CallerMust()
	fmt.Printf("%s\n", gdk.JsonPretty(c))
	fmt.Printf("%v\n", c.Path)

	sys.Organize = "gdk"
	c = sys.CallerMust()
	fmt.Printf("%s", gdk.JsonPretty(c))

	// Output
	// {
	//	"PC": 17790296,
	//	"Path": "/Users/xxx/github/gdk/sys/fn_test.go",
	//	"Name": "github.com/hollson/gdk/sys_test.ExampleCallerMust",
	//	"Line": 19
	// }
}
