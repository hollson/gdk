// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package emoji

// Emoji Emoji类型
//  参考：https://www.unicode.org/Public/emoji/15.0/emoji-sequences.txt
type Emoji rune

const (
	Info  Emoji = 0x1F514 // 铃铛 🔔
	Warn  Emoji = 0x26a0  // 警告 ⚠️
	Error Emoji = 0x274C  // 错误 ❌
	Ok    Emoji = 0x2705  // 成功 ✅
	No    Emoji = 0x1f6ab // 失败 🚫
)

func (e Emoji) String() string {
	return string(e)
}

// Render 将Unicode格式的Emoji字符渲染成对应的图标
//  参考：https://www.unicode.org/Public/emoji/15.0/emoji-sequences.txt
func Render(r rune) string {
	return string(r)
}
