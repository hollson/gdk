// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"testing"
)

func TestNumberString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandNumStr(20))
	}
}

func TestLetterString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandString(20))
	}
}

func TestPasswordString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandPassword(20))
	}
}

func TestRandSequence(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandSequence(5))
	}
}

func TestRandScopeInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandScopeInt(0, 10))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(RandScopeInt(100, 999))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(RandScopeInt(10000, 99999))
	}
}

func TestRandInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandInt64(1000))
	}
}

func TestRandInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandInt(1000))
	}
}
