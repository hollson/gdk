// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package memory

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	fmt.Println(Size(0.5))                       // 0.50B
	fmt.Println(Size(128.88))                    // 128.88B
	fmt.Println(Size(1024 + 512 + 256))          // 1.75KB
	fmt.Println((Size(85509571)).String())       // 81.55MB
	fmt.Println((MB * 0.5).String())             // 333 512.00KB
	fmt.Println((MB * 32).String())              // 32.00MB
	fmt.Println((MB * 1024).String())            // 1.00GB
	fmt.Println((MB * 2096).String())            // 2.05GB
	fmt.Println((GB * 32).String())              // 32.00GB
	fmt.Println((GB * 2 * 1024 * 1024).String()) // 2.00PB
	fmt.Println((EB * 64).String())              // 64.00EB
	fmt.Println((ZB * (1024 + 128)).String())    // 1.12YB
	fmt.Println((YB * 0.5).String())             // 512.00ZB
}
