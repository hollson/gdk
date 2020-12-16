package util

import (
    "strconv"
)

// 数据转换
func StrToInt64(num string) int64 {
    if nt64, ie := strconv.ParseInt(num, 10, 64); ie != nil {
        return 0
    } else {
        return nt64
    }
}



// 三目运算
func T3(b bool, x, y interface{}) interface{} {
    if b {
        return x
    }
    return y
}
