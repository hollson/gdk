// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"testing"
)

func TestMatchAny(t *testing.T) {
	fmt.Println(MatchAny("golang", "foobarbaz", "hello world", "golang")) // true
	fmt.Println(MatchAny("go", "foobarbaz", "hello world", "golang"))     // false
	fmt.Println(MatchAny("go*", "foobarbaz", "hello world", "golang"))    // true

	fmt.Println(MatchAny("foo*", "foobarbaz", "hello world", "golang"))      // true
	fmt.Println(MatchAny("*lang", "foobarbaz", "hello world", "golang"))     // true
	fmt.Println(MatchAny("foo*baz", "foobarbaz", "hello world", "golang"))   // true
	fmt.Println(MatchAny("foo*ba", "foobarbaz", "hello world", "golang"))    // false
	fmt.Println(MatchAny("foo*lang", "foobarbaz", "hello world", "golang"))  // false
	fmt.Println(MatchAny("&^%$*lang", "foobarbaz", "hello world", "golang")) // false
}
