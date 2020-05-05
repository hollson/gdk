//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-01-20
// @ Version: 1.0.0
//
// 元素是否包含在某个集合中
//-------------------------------------------------------------------------------------

package utility


// byte
func ContainByte(tar byte, all ...byte) bool {
	for _, v := range all {
		if v == tar {
			return true
		}
	}
	return false
}

// int
func ContainInt(tar int, all ...int) bool {
	for _, v := range all {
		if v == tar {
			return true
		}
	}
	return false
}

// int64
func ContainInt64(tar int64, all ...int64) bool {
	for _, v := range all {
		if v == tar {
			return true
		}
	}
	return false
}

// float32
func ContainFloat32(tar float32, all ...float32) bool {
	for _, v := range all {
		if v == tar {
			return true
		}
	}
	return false
}

//float64
func ContainFloat64(tar float64, all ...float64) bool {
	for _, v := range all {
		if v == tar {
			return true
		}
	}
	return false
}

//string
func ContainStr(tar string, all ...string) bool {
	for _, v := range all {
		if tar == v {
			return true
		}
	}
	return false
}