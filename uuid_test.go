package gdk

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(UUID())
	}

	fmt.Println()
	fmt.Println(UUIDShort())
}
