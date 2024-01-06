// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sys_test

import (
	"fmt"
	"testing"

	"github.com/hollson/gdk/sys"
)

func TestErrorf(t *testing.T) {
	ExampleErrorf()

	sys.DefaultErrFormat = func(c *sys.CallerInfo, format string, a ...interface{}) string {
		return fmt.Sprintf("🔴 %s:%d %s ", c.Name, c.Line, fmt.Sprintf(format, a...))
	}
	fmt.Println(sys.Errorf("%s", "一个错误🙅"))
}

func ExampleErrorf() {
	sys.Organize = "gdk"
	fmt.Println(sys.Errorf("===%s===", "sorry"))

	// Output:
	// ❌ gdk/sys_test.ExampleErrorf:25 ===[sorry]===
}
