// Package color 精简快捷的彩色字体工具包(Linux版本)
package color

// StyleOrColor 样式/颜色
type StyleOrColor uint8

const (
	Reset        StyleOrColor = 0 // 重置
	Bold                      = 1 // 加粗
	Faint                     = 2 // 模糊
	Italic                    = 3 // 斜体
	Underline                 = 4 // 下划线
	BlinkSlow                 = 5 // 慢速闪烁
	BlinkRapid                = 6 // 快速闪烁
	ReverseVideo              = 7 // 反白/反向显示
	Concealed                 = 8 // 隐藏/暗格
	CrossedOut                = 9 // 删除

	FontBlack   = 30 // 「字体」黑色
	FontRed     = 31 // 「字体」红色
	FontGreen   = 32 // 「字体」绿色
	FontYellow  = 33 // 「字体」黄色
	FontBlue    = 34 // 「字体」蓝色
	FontMagenta = 35 // 「字体」品红/洋紫
	FontCyan    = 36 // 「字体」青色
	FontWhite   = 37 // 「字体」白色

	BackBlack   = 40 // 「背景」黑色
	BackRed     = 41 // 「背景」红色
	BackGreen   = 42 // 「背景」绿色
	BackYellow  = 43 // 「背景」黄色
	BackBlue    = 44 // 「背景」蓝色
	BackMagenta = 45 // 「背景」品红/洋紫
	BackCyan    = 46 // 「背景」青色
	BackWhite   = 47 // 「背景」白色
)
