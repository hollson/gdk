// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"testing"
)

func TestConvert2AnyTypeSlice(t *testing.T) {
	var arr = []int{1, 2, 3}
	fmt.Println(IsSlice(arr))
	fmt.Println(Convert2AnyTypeSlice(arr))
	fmt.Println(Convert2AnyTypeSlice("abc"))
	fmt.Println(Convert2AnyTypeSlice(func() {}))
	fmt.Println(Convert2AnyTypeSlice(map[int]string{}))
}

func TestContains(t *testing.T) {
	fmt.Println(ContainByte('a', 97, 98, 99))
	fmt.Println(ContainInt(1, 1, 2, 3))
	fmt.Println(ContainInt32(int32(1), 1, 2, 3))
	fmt.Println(ContainInt64(int64(1), 1, 2, 3))
	fmt.Println(ContainFloat32(1.2, 1.1, 2.2, 3.3))
	fmt.Println(ContainFloat64(1.1, 1.1, 2.2, 3.3))
	fmt.Println(ContainString("a", "a", "b", "c"))
	fmt.Println(ContainInt(100, []int{1, 2, 3}...))
}

func TestContain(t *testing.T) {
	type User struct{ Name string }
	var ints = []int{1, 2, 3}
	var a, b, c, d = 1, 2, 3, 3

	fmt.Println(Contain(1, ints))
	fmt.Println(Contain(20, ints))
	fmt.Println(Contain(1, []int{1, 2, 3}))
	fmt.Println(Contain(&a, []*int{&a, &b, &c}))
	fmt.Println(Contain(&d, []*int{&a, &b, &c}))
	fmt.Println(Contain(1, []string{"a", "b", "c"}))
	fmt.Println(Contain(1, "hello"))
	fmt.Println(Contain(1.2, []interface{}{1.2, "a", 'b'}))

	var us = []*User{{Name: "AA"}, {Name: "BB"}, {Name: "CC"}}
	fmt.Println(Contain(User{Name: "AA"}, us))
	fmt.Println(Contain(&User{Name: "AA"}, us))
	fmt.Println(Contain(us[0], us))
}

func TestContainMust(t *testing.T) {
	fmt.Println(ContainMust(1, []int{1, 2, 3}))
	fmt.Println(ContainMust(1, "abc"))
}

// func TestDistinct(t *testing.T) {
// 	var arr = []int{1, 2, 2, 3, 3, 3}
// 	fmt.Println(Distinct(arr))
// }
