// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package emoji

type Emoji rune

const (
	Info  Emoji = 0x1F514 // 🔔
	Warn  Emoji = 0x26a0  // ⚠️
	Error Emoji = 0x274C  // ❌
	Ok    Emoji = 0x2705  // ✅
	No    Emoji = 0x1f6ab // 🚫
)

func (e Emoji) String() string {
	return string(e)
}

func Convert(r rune) string {
	return string(r)
}
