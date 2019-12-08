package utilx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 切片是否包含目标元素
func ContainInt(one int, all ...int) bool {
	for _, v := range all {
		if v == one {
			return true
		}
	}
	return false
}

// 切片是否包含目标元素
func ContainStr(one string, all ...string) bool {
	for _, v := range all {
		if one == v {
			return true
		}
	}
	return false
}

// 切片去重
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i - 1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i))
	}
	return ret
}

// 数据转换
func StrToInt64(num string) int64 {
	if nt64, ie := strconv.ParseInt(num, 10, 64); ie != nil {
		return 0
	} else {
		return nt64
	}
}



//将字符串数组拼接成sqlin
func ToSqlIn(items ...string) string {
	return fmt.Sprintf(`(%s)`, T3(len(items) > 0, strings.Join(items, ","), "''"))
}


// 三目运算
func T3(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}









