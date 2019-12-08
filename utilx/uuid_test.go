//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-11-28
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

package utilx

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(Guid())
	}
}