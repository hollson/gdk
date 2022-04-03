// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package memory

import (
	"fmt"
)

// Size 内存大小(单位：字节)
//  Reference：https://golang.org/doc/effective_go#constants
type Size float64

const (
	_       = iota
	KB Size = 1 << (10 * iota) // 千字节
	MB                         // 兆字节
	GB                         // 吉字节/千兆
	TB                         // 太字节
	PB                         // 拍字节
	EB                         // 艾字节
	ZB                         // 泽字节
	YB                         // 尧字节
)

func (b Size) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
