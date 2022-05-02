// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package money

import (
	"fmt"
	"testing"
)

func TestCNY_String(t *testing.T) {
	// 构造
	fmt.Println(New(0, 0, 9))
	fmt.Println(New(0, 9, 9))
	fmt.Println(New(1, 9, 9))
	fmt.Println(New(100, 9, 9))
	fmt.Println(New(1000, 9, 9))
	fmt.Println(New(10000, 0, 0))
	fmt.Println(New(1111_2222, 9, 9))
	fmt.Println(New(1_2222_3333, 0, 0))

	fmt.Println(2 * W)
	fmt.Println(9 * F)
	fmt.Println(9*J + 9*F)
	fmt.Println(456*Y + 9*J + 9*F)
	fmt.Println(3456*Y + 9*J + 9*F)
	fmt.Println(2222*W + 3456*Y + 9*J + 9*F)
	fmt.Println(1*E + 2222*W + 3456*Y + 9*J + 9*F)

	fmt.Println(CNY(1289))
}
