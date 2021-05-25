// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"fmt"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	fmt.Printf("[%s]\n", TrimSpace(`
	select *   
  from  \n\t
wherw \n 1=1
order by \t\t\t\ id desc       `))

	fmt.Printf("[%s]\n", TrimSpace("select \n * from \t\r student  where 1=1        \n"))
}
