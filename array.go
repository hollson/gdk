// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"reflect"
)

// 判断参数o是否为某个类型切片
//  如: IsSlice([]int{1,2,3}) => true
func IsSlice(o interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(o)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

// 将对象切片转换为接口切片
//  如: AnyTypeSlice([]int{1,2,3}) => []interface{}{1,2,3}
func ObjectSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := IsSlice(slice)
	if !ok {
		return nil, false
	}

	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, true
}
