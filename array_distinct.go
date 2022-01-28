// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

// 目标元素tar是否包含在container集合中
func DistinctBytes(tar byte, container ...byte) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func DistinctInts(tar int, container ...int) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func DistinctInt32s(tar int32, container ...int32) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func DistinctInt64s(tar int64, container ...int64) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func DistinctFloat32s(tar float32, container ...float32) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func DistinctFloat64s(tar float64, container ...float64) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func DistinctStrings(tar string, container ...string) bool {
	for _, v := range container {
		if tar == v {
			return true
		}
	}
	return false
}

//
// // 切片去重 Distinct
// func Distinct(a interface{}) ([]interface{}, error) {
// 	// arr, err := Convert2AnyTypeSlice(a)
// 	if err != nil {
//
// 	}
// 	var ret []interface{}
//
// 	va := reflect.ValueOf(a)
// 	for i := 0; i < va.Len(); i++ {
// 		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
// 			continue
// 		}
// 		ret = append(ret, va.Index(i))
// 	}
// 	return ret
// }

func distinctStringArray(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// https://blog.csdn.net/qq_27068845/article/details/77407358
