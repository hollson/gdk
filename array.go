// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"errors"
	"fmt"
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

// Convert2AnyTypeSlice 将对象切片转换为接口切片
//  如: AnyTypeSlice([]int{1,2,3}) => []interface{}{1,2,3}
func Convert2AnyTypeSlice(slice interface{}) ([]interface{}, error) {
	val, ok := IsSlice(slice)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%#v is not a slice type", reflect.ValueOf(slice)))
	}

	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, nil
}


