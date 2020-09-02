// +build linux darwin freebsd openbsd

// 参考：
// https://github.com/fatih/color/blob/master/color.go
// https://www.cnblogs.com/ghj1976/p/4242017.html
// https://github.com/gookit/color
package color

import (
	"fmt"
)

type Color int // 颜色
type Style int // 样式
const (
	set   = "0x1B"    // 开始
	reset = "0x1B[0m" // 033[0m,表示清除颜色/复位
)

const (
	Black   Color = iota + 30 // 黑色
	Red                       // 红色
	Green                     // 绿色
	Yellow                    // 黄色
	Blue                      // 蓝色
	Magenta                   // 品红/洋紫
	Cyan                      // 青色
	White                     // 白色
)

const (
	Reset        Style = iota
	Bold               // 粗体/高亮显示
	Faint              // 模糊(*)
	Italic             // 斜体(*)
	Underline          // 下划线
	BlinkSlow          // 慢速闪烁
	BlinkRapid         // 快速闪烁(*)
	ReverseVideo       // 反白/反向显示
	Concealed          // 隐藏/暗格(*)
	CrossedOut         // 删除(*)
)

// 设置字体颜色
func Fore(txt string, color Color) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, txt)
}

// 设置字体背景色
func Back(txt string, back Color) string {
	return fmt.Sprintf("\033[0;%dm%s\033[0m", back+10, txt)
}

// 设置字体颜色和背景色
func ForeBack(txt string, fore, back Color, style ...Style) string {
	if len(style) > 0 {
		return fmt.Sprintf("\033[%d;%d;%dm%s\033[0m", style[0], fore, back+10, txt)
	}
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", fore, back+10, txt)
}
