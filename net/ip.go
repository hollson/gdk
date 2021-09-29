// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package net

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// https://blog.csdn.net/chinawangfei/article/details/89965515
// 局域网IP
func LocalIP() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, address := range as {
		// 排除回环地址
		if _ip, ok := address.(*net.IPNet); ok && !_ip.IP.IsLoopback() && _ip.IP.To4() != nil {
			return _ip.IP, nil
		}
	}
	return nil, errors.New("not found available ip")
}

// 公网IP
func PubIp() string {
	loc, zone := time.Now().Zone()
	if zone/3600 == 8 {
		fmt.Println(loc)
		// http.Get("")
	}
	fmt.Println(exec.Command("curl","--connect-timeout 1 -m 3 ip.cip.cc").String())

	pong, err := http.Get("http://ip.cip.cc/")
	if err != nil {
		fmt.Println("===>", err.Error())
		return ""
	}
	data, err := io.ReadAll(pong.Body)
	if err != nil {
		fmt.Println("===>", err.Error())
		return ""
	}
	fmt.Println(string(data))
	return string(data)
}

// 公网IP
// curl --connect-timeout 1 -m 3 ifconfig.me
// curl icanhazip.com	美国
// curl ip.sb		美国
// curl ifconfig.io	美国
// curl httpbin.org/ip
// curl cip.cc     汕头
// curl --connect-timeout 1 -m 3 ip.cip.cc

// 当前主机的局域网IP
func IPByDNS() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {

	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
