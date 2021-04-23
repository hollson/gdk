// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package money

import (
    "fmt"
    "testing"
)

func TestCNY_Float64(t *testing.T) {
    fmt.Println((123456.123456789876 * E).Float64())
}

func TestCNY_String(t *testing.T) {
    fmt.Println(CNY(8.01))
    fmt.Println(CNY(12.3))
    fmt.Println(CNY(1289))
    fmt.Println(1234.9 * J)
    fmt.Println(1.23 * Q)
    fmt.Println(44.2 * Q)
    fmt.Println(4.43 * W)
    fmt.Println(4.568 * W)
    fmt.Println(1200.99 * W)
    fmt.Println(1200.888 * E)
    fmt.Println(12345.777 * E)
    fmt.Println(123456789.12 * E)
}

func TestCNY_Single(t *testing.T) {
    fmt.Println(CNY(8.01).Single())
    fmt.Println(CNY(12.3).Single())
    fmt.Println(CNY(1289).Single())
    fmt.Println((1234.9 * J).Single())
    fmt.Println((1.23 * Q).Single())
    fmt.Println((44.2 * Q).Single())
    fmt.Println((4.43 * W).Single())
    fmt.Println((4.568 * W).Single())
    fmt.Println((1200.99 * W).Single())
    fmt.Println((1200.888 * E).Single())
    fmt.Println((12345.777 * E).Single())
    fmt.Println((123456789.12 * E).Single())
}
