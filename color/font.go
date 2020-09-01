// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package color

import "fmt"

type FontColor int



const (
	Black  FontColor = iota + 30 // 黑色
	Red                          // 红色
	Green                        // 绿色
	Yellow                       // 黄色
	Blue                         // 蓝色
	Purple                       // 紫色
	Cyan                         // 青色
	White                        // 白色
)

type Style int

const (
	Default   Style = 0 // 终端默认设置
	Highlight       = 1 // 高亮显示
	Underline       = 4 // 使用下划线
	Glitter         = 5 // 闪烁
	Reverse         = 7 // 反白显示
	Invisible       = 8 // 不可见
)

// 设置字体颜色
func Fore(txt string, color FontColor) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, txt)
}

// 设置字体背景色
func Back(txt string, back FontColor) string {
	return fmt.Sprintf("\033[0;%dm%s\033[0m", back+10, txt)
}

// 设置字体颜色和背景色
func ForeBack(txt string, fore, back FontColor, style ...Style) string {
	if len(style) > 0 {
		return fmt.Sprintf("\033[%d;%d;%dm%s\033[0m", style[0], fore, back+10, txt)
	}
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", fore, back+10, txt)
}

// 使用示例：
// func main() {
// 	fmt.Println(Fore("字体颜色", Purple))
// 	fmt.Println(Back("背景颜色", Green))
// 	fmt.Println(ForeBack("字体和背景", Blue, Yellow))
// 	fmt.Println(ForeBack("字体和背景&样式", Red, Cyan, Underline))
// }