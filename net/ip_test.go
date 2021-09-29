// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package net

import (
	"fmt"
	"testing"
)
// //https://www.zeitverschiebung.net/cn/city/2643743
// func TestLocalIP(t *testing.T) {
// 	loc, offset := time.Now().Zone()
// 	fmt.Println("本地：", loc, offset/3600)
//
// 	var cstSh, _ = time.LoadLocation("Asia/Shanghai") // 上海
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("上海：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("Asia/Taipei")
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("台北：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("Asia/Seoul")
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("首尔：", loc, offset/3600)
//
//
// 	cstSh, _ = time.LoadLocation("Asia/Tokyo")
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("东京：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("Europe/London")
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("伦敦：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("America/Los_Angeles") // 美国
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("洛杉矶：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("America/New_York") // 美国
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("纽约：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("Asia/Jerusalem") // 美国
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("耶路撒冷：", loc, offset/3600)
//
// 	cstSh, _ = time.LoadLocation("Europe/Dublin") // 美国
// 	loc, offset=time.Now().In(cstSh).Zone()
// 	fmt.Println("都柏林：", loc, offset/3600)
//
// 	// fmt.Println("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
//
// 	// d,_:=net.InterfaceAddrs()
// 	// fmt.Printf("%+v",d)
// 	fmt.Println(LocalIP())
// 	// fmt.Println( LocalIP())
// }

func TestPubIp(t *testing.T) {
	fmt.Println(PubIp())
}


//https://zhuanlan.zhihu.com/p/296409942