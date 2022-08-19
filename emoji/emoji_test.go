// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package emoji_test

import (
	"fmt"
	"testing"

	"github.com/hollson/gdk/emoji"
)

func TestEmoji_String(t *testing.T) {
	fmt.Println(1, emoji.Info)
	fmt.Println(2, emoji.Warn)
	fmt.Println(3, emoji.Error)
	fmt.Println(4, emoji.Ok)
	fmt.Println(5, emoji.No)

	fmt.Println(6, emoji.Render(0x1F514))
	fmt.Println(7, emoji.Render(0x274C))
	fmt.Println(8, emoji.Render(0x2705))
	fmt.Printf("%0x\n", '🚫')
}

func ExampleRender() {
	fmt.Println(emoji.Render(0x1F514))
	fmt.Println(emoji.Render(0x26a0))
	fmt.Println(emoji.Render(0x274C))
	fmt.Println(emoji.Render(0x2705))
	fmt.Println(emoji.Render(0x1f6ab))

	// output：
	// 	🔔
	// 	⚠
	// 	❌
	// 	✅
	// 	🚫
}

func TestRender(t *testing.T) {
	ExampleRender()
}
