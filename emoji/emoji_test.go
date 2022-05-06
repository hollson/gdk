// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package emoji_test

import (
	"fmt"
	"testing"

	"github.com/hollson/gdk/emoji"
)

func TestToEmoji(t *testing.T) {
	fmt.Println(1, emoji.Info)
	fmt.Println(2, emoji.Warn)
	fmt.Println(3, emoji.Error)
	fmt.Println(4, emoji.Ok)
	fmt.Println(5, emoji.No)

	fmt.Println(6, emoji.Emoji(0x1F514))
	fmt.Println(7, emoji.Emoji(0x274C))
	fmt.Println(8, emoji.Convert(0x2705))
	fmt.Printf("%0x\n", '🚫')
}
