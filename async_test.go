// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"testing"
	"time"
)

func ExampleGo() {
	ExceptionHandler = func(err interface{}) {
		fmt.Printf("custom error: %v\n", err)
	}

	Go(func() {
		panic("just a test")
	})

	time.Sleep(time.Second)

	// Output:
	// custom error: just a test
}

func TestGo(t *testing.T) {
	ExampleGo()
}
