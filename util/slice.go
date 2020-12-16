package util

import (
    "fmt"
    "reflect"
    "strings"
)

// 切片反转 - byte
func ReverseByteSlice(s []byte) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片反转 - int
func ReverseIntSlice(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片反转 - int32
func ReverseInt32Slice(s []int32) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片反转 - int64
func ReverseInt64Slice(s []int64) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片反转 - string
func ReverseStringSlice(s []string) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片反转 - float32
func ReverseFloat32Slice(s []float32) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片反转 - float64
func ReverseFloat64Slice(s []float64) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 切片去重
func Duplicate(a interface{}) (ret []interface{}) {
    va := reflect.ValueOf(a)
    for i := 0; i < va.Len(); i++ {
        if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
            continue
        }
        ret = append(ret, va.Index(i))
    }
    return ret
}

// 将切片拼接成「sql in」语句
//  如：{"aa","bb","cc"} ==> "'1','3','5','7','9'"
func SQLInString(items ...string) (inSql string) {
    if len(items) == 0 {
        return "''"
    }
    for _, v := range items {
        inSql = fmt.Sprintf("%s,'%s'", inSql, v)
    }
    return strings.TrimLeft(inSql, ",")
}

// 将切片拼接成「sql in」语句
//  如：{1,3,5,7,9} ==> "1,3,5,7,9"
func SQLInInt(items ...int) (inSql string) {
    if len(items) == 0 {
        return "''"
    }
    for _, v := range items {
        inSql = fmt.Sprintf("%s,%d", inSql, v)
    }
    return strings.TrimLeft(inSql, ",")
}

// 将切片拼接成「sql in」语句
//  如：{1,3,5,7,9} ==> "1,3,5,7,9"
func SQLInInt32(items ...int32) (inSql string) {
    if len(items) == 0 {
        return "''"
    }
    for _, v := range items {
        inSql = fmt.Sprintf("%s,%d", inSql, v)
    }
    return strings.TrimLeft(inSql, ",")
}

// 将切片拼接成「sql in」语句
//  如：{1,3,5,7,9} ==> "1,3,5,7,9"
func SQLInInt64(items ...int64) (inSql string) {
    if len(items) == 0 {
        return "''"
    }
    for _, v := range items {
        inSql = fmt.Sprintf("%s,%d", inSql, v)
    }
    return strings.TrimLeft(inSql, ",")
}
