// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-05-06
// @ Version: 1.0.0
//
// Here's the code description...
// -------------------------------------------------------------------------------------

package memory

import (
	"fmt"
)

// 存储容量
type Size int

const (
	B  Size = 1 << (10 * iota) // 1 << (10*0)
	KB                         // 1 << (10*1)
	MB                         // 1 << (10*2)
	GB                         // 1 << (10*3)
	TB                         // 1 << (10*4)
	PB                         // 1 << (10*5)
	EB                         // 1 << (10*6)
)

func (b Size) String() string {
	switch {
	case b >= EB:
		return fmt.Sprintf("%.2fEB", float64(b)/float64(EB))
	case b >= PB:
		return fmt.Sprintf("%.2fPB", float64(b)/float64(PB))
	case b >= TB:
		return fmt.Sprintf("%.2fTB", float64(b)/float64(TB))
	case b >= GB:
		return fmt.Sprintf("%.2fGB", float64(b)/float64(GB))
	case b >= MB:
		return fmt.Sprintf("%.2fMB", float64(b)/float64(MB))
	case b >= KB:
		return fmt.Sprintf("%.2fKB", float64(b)/float64(KB))
	default:
		return fmt.Sprintf("%.2fB", float64(b))
	}
}
