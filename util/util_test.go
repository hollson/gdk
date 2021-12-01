// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
	"time"
)

func ExampleZone() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	name, zone := Zone(loc)
	fmt.Println(name, zone)

	// out: CST +8
}

func ExampleShiftZone() {
	// 将当前时间转换为纽约时间
	loc, _ := time.LoadLocation("America/New_York")
	_newYorkTime := ShiftZone(loc, time.Now())
	_newYorkTime = ShiftZone(loc) // 默认即为Nows
	fmt.Println(_newYorkTime)

	// out: 2006-01-02 13:04:05.868597 -0500 EST
}

func TestLocalIP(t *testing.T) {
	var zs = map[string]string{
		"Asia/Shanghai":       "上海",
		"Asia/Taipei":         "台北",
		"Asia/Seoul":          "首尔",
		"Asia/Tokyo":          "东京",
		"Asia/Singapore":      "新加坡",
		"Asia/Jerusalem":      "耶路撒冷",
		"Asia/Kolkata":        "加尔各答", // 印度
		"America/Los_Angeles": "洛杉矶",
		"America/New_York":    "纽约",
		"Europe/Dublin":       "都柏林",
		"Europe/London":       "伦敦",
	}

	for k, v := range zs {
		var loc, err = time.LoadLocation(k)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		name, zone := Zone(loc)
		fmt.Printf("%v  %v  %v\n", v, name, zone)
	}
}

func TestToMonday(t *testing.T) {
	fmt.Println(ToMonday(time.Now().Add(-time.Hour * 24)))
}

func TestShiftZone(t *testing.T) {
	ExampleShiftZone()
}
