// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package host

import (
    "net"
    "strings"
)

// https://blog.csdn.net/chinawangfei/article/details/89965515
func foo() (addr []string) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return []string{"0.0.0.0"}
    }
    for _, address := range addrs {
        // 检查ip地址判断是否回环地址
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                addr = append(addr, ipnet.IP.String())
            }
        }
    }
    return
}

func LocalIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {

    }
    defer conn.Close()
    localAddr := conn.LocalAddr().String()
    idx := strings.LastIndex(localAddr, ":")
    return localAddr[0:idx]
}
