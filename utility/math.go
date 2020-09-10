//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-11-29
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

package utility

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandIntn(n int) int {
	// 创建随机数种子  (默认将1970-1-1 00:00:00当做随机数种子)
	rand.Seed(time.Now().UnixNano()) // UnixNano()表示纳秒
	a := rand.Intn(n)               // 0-n之间的随机整数。 (包含0，不包含n。相当于对10求余)
	return a
}

// 保留两位
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//https://blog.csdn.net/zhizhengguan/article/details/93380945


//golang 产生不重复的随机数
//package utils
//import (
//"math/rand"
//"time"
//)
//
//var channel chan int64 = make(chan int64, 32)
//
//func init() {
//	go func() {
//		var old int64
//		for {
//			o := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
//			if old != o {
//				old = o
//				select {
//				case channel <- o:
//				}
//			}
//		}
//	}()
//}
//func RandInt64() (r int64) {
//	select {
//	case rand := <-channel:
//		r = rand
//	}
//	return
//}
