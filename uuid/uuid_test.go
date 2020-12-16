// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid

import (
    "fmt"
    "testing"
)

func TestUUID(t *testing.T) {
    for i := 0; i < 10; i++ {
        fmt.Println(NewV1())
    }
    fmt.Println()
    for i := 0; i < 10; i++ {
        fmt.Println(NewV2(127))
    }
    fmt.Println()
    for i := 0; i < 10; i++ {
        // v4, _ := NewV1()
        fmt.Println(NewV3(Nil, ""))
    }
    fmt.Println()
    for i := 0; i < 10; i++ {
        fmt.Println(NewV4())
    }
    fmt.Println()
    for i := 0; i < 10; i++ {
        // v4, _ := NewV1()
        fmt.Println(NewV5(Nil, "v5"))
    }

    fmt.Println()
    for i := 0; i < 10; i++ {
        // v4, _ := NewV1()
        fmt.Println(New())
    }
}
