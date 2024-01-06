// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

// StringsToInts string slice to int slice
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

// StringsToSlice convert []string to []interface{}
func StringsToSlice(strings []string) []interface{} {
	args := make([]interface{}, len(strings))
	for i, s := range strings {
		args[i] = s
	}
	return args
}

// StringsRemove a value form a string slice
func StringsRemove(ss []string, s string) []string {
	ns := make([]string, 0, len(ss))
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}
	return ns
}

// TrimStrings trim string slice item.
//
// Usage:
//	// output: [a, b, c]
// 	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
func TrimStrings(ss []string, cutSet ...string) (ns []string) {
	cutSetLn := len(cutSet)
	hasCutSet := cutSetLn > 0 && cutSet[0] != ""

	var trimSet string
	if hasCutSet {
		trimSet = cutSet[0]
	}
	if cutSetLn > 1 {
		trimSet = strings.Join(cutSet, "")
	}

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, trimSet))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}
	return
}

// GetRandomOne get random element from an array/slice
func GetRandomOne(arr interface{}) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return arr
	}

	i := RandScopeInt(0, rv.Len())
	r := rv.Index(i).Interface()
	return r
}
