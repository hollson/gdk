// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package netx

import (
	"fmt"
	"testing"
)

func TestInnerIP4(t *testing.T) {
	fmt.Println(InnerIP4())
}

func TestOuterIPWithProxy(t *testing.T) {
	for k, v := range _proxy {
		ip, err := OuterIPWithProxy(v)
		if err != nil {
			fmt.Printf("Sorry：%v,%v\n", v, err)
			continue
		}
		fmt.Printf("%d「%v」 - %v\n", k, v, ip.String())
	}
}

func TestOuterIP(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(OuterIP())
	}
}

func TestInnerIPWithDNS(t *testing.T) {
	fmt.Println(InnerIPWithDNS())
}

// 推荐文章： https://blog.csdn.net/chinawangfei/article/details/89965515
