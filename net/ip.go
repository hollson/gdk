// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package net

import (
	"errors"
	"net"
	"strings"
)

// https://blog.csdn.net/chinawangfei/article/details/89965515
func IP() (net.IP, error) {
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
