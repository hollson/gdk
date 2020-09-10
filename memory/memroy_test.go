// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package memory

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	fmt.Println((Size(32)).String())
	fmt.Println((MB * 32).String())
	fmt.Println((MB * 1024).String())
	fmt.Println((MB * 1025).String())
	fmt.Println((GB * 32).String())
	// fmt.Println((EB * 64).String())
}
