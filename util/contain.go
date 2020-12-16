// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-01-20
// @ Version: 1.0.0
//
// 元素是否包含在某个集合中
// -------------------------------------------------------------------------------------

package util

// 元素o是否包含在切片a中
//  byte
func ContainByte(o byte, a ...byte) bool {
    for _, v := range a {
        if v == o {
            return true
        }
    }
    return false
}

// 元素o是否包含在切片a中
//  int
func ContainInt(o int, a ...int) bool {
    for _, v := range a {
        if v == o {
            return true
        }
    }
    return false
}

// 元素o是否包含在切片a中
//  int32
func ContainInt32(o int32, a ...int32) bool {
    for _, v := range a {
        if v == o {
            return true
        }
    }
    return false
}

// 元素o是否包含在切片a中
//  int64
func ContainInt64(o int64, a ...int64) bool {
    for _, v := range a {
        if v == o {
            return true
        }
    }
    return false
}

// 元素o是否包含在切片a中
//  float32
func ContainFloat32(o float32, a ...float32) bool {
    for _, v := range a {
        if v == o {
            return true
        }
    }
    return false
}

// 元素o是否包含在切片a中
//  float64
func ContainFloat64(o float64, a ...float64) bool {
    for _, v := range a {
        if v == o {
            return true
        }
    }
    return false
}

// 元素o是否包含在切片a中
//  string
func ContainString(o string, a ...string) bool {
    for _, v := range a {
        if o == v {
            return true
        }
    }
    return false
}
