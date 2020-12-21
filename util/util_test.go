// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
    "fmt"
    "testing"
)

func TestContainFloat64(t *testing.T) {
    var a int = 12
    fmt.Println(ContainInt(a, 12, 33, 55))
}

func TestSQLIn(t *testing.T) {
    var a = []int64{1, 3, 5, 7, 9}
    s := SQLInInt64(a...)
    fmt.Printf("「%s」\n", s)

    b := []string{"aa", "bb", "cc"}
    s = SQLInString(b...)
    fmt.Printf("「%s」\n", s)
}

func TestJsonFormat(t *testing.T) {
    person := struct {
        Name string
        Age  int
    }{"tom", 22}
    fmt.Println(JsonFormat(person))
}
