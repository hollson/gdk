// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package netx

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	_innerIP4 net.IP
	_outerIP4 net.IP
	_proxy    = map[int]string{
		0: "http://ip.cip.cc",
		1: "http://icanhazip.com",
		2: "http://ifconfig.me",
		3: "http://ifconfig.io/ip",
		4: "https://api.ip.sb/ip",
		5: "http://ipinfo.io/ip",
		6: "http://ipecho.net/plain",
	}
)

// InnerIP4 获取局域网IP4地址
func InnerIP4() (net.IP, error) {
	if _innerIP4 != nil {
		return _innerIP4, nil
	}
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, address := range as {
		if _ip, ok := address.(*net.IPNet); ok && !_ip.IP.IsLoopback() && _ip.IP.To4() != nil {
			return _ip.IP, nil
		}
	}
	return nil, errors.New("未找到可用的内网IP4地址")
}

// InnerIPWithDNS 使用在线DNS查询内网IP
func InnerIPWithDNS() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ep := conn.LocalAddr().String()
	idx := strings.IndexRune(ep, ':')
	if idx < 1 {
		return nil, errors.New("failed to get ip using DNS")
	}
	return net.ParseIP(ep[0:idx]), nil
}

func OuterIP() (net.IP, error) {
	if _outerIP4 != nil {
		return _outerIP4, nil
	}
	for _, v := range _proxy {
		ret, _err := OuterIPWithProxy(v)
		if _err == nil {
			_outerIP4 = ret
			return ret, nil
		}
		fmt.Printf("IP query error:  %v\n", v)
	}
	return nil, errors.New("get outer IP failed")
}

// OuterIPWithProxy 根据第三方IP服务,获取当前网络的公网IP
func OuterIPWithProxy(proxyUrl string) (net.IP, error) {
	cli := http.Client{Timeout: time.Second * 3}
	req, err := http.NewRequest("GET", proxyUrl, nil)
	if err != nil {
		return nil, err
	}
	// req.Header.Add("User Agent","curl/7.64.1")
	req.Header.Add("Accept", "application/json,text/plain,*/*;q=0.01")

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("data error: %v, err:%v", data, err)
	}

	ip := net.ParseIP(strings.TrimRight(string(data), "\n"))
	if ip == nil {
		return nil, fmt.Errorf("query ip failed, %v", string(data))
	}
	return ip, nil
}

// 将IP转换成int32或int64类型
func IPCode(ip net.IP) int {
	return 1
}
