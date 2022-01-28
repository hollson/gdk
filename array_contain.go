// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"errors"
	"reflect"
)

// 目标元素tar是否包含在container集合中
func ContainByte(tar byte, container ...byte) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func ContainInt(tar int, container ...int) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func ContainInt32(tar int32, container ...int32) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func ContainInt64(tar int64, container ...int64) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func ContainFloat32(tar float32, container ...float32) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func ContainFloat64(tar float64, container ...float64) bool {
	for _, v := range container {
		if v == tar {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中
func ContainString(tar string, container ...string) bool {
	for _, v := range container {
		if tar == v {
			return true
		}
	}
	return false
}

// 目标元素tar是否包含在container集合中,container必须是切片类型
//  如：Contain(1,[]int{1,2,3})
//     Contain(u,[]User{u1,u2,u3})
//     Contain(&u,[]*User{&u1,&u2,&u3})
//     Contain(1.0, []interface{}{1.0,"a",'b'})
func Contain(tar interface{}, container interface{}) (bool, error) {
	if arr, err := Convert2AnyTypeSlice(container); err!=nil {
		return contain(tar, arr), nil
	}
	return false, errors.New("container is not a slice type")
}

// 同Contain，但需手动处理panic
func ContainMust(tar interface{}, container interface{}) bool {
	b, err := Contain(tar, container)
	if err != nil {
		panic(err)
	}
	return b
}

func contain(tar interface{}, container []interface{}) bool {
	a := reflect.ValueOf(tar)
	for _, v := range container {
		if a.Interface() == v {
			return true
		}
		// reflect.DeepEqual(a.Interface(), reflect.ValueOf(v))
	}
	return false
}
