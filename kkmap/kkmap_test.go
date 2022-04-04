// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package kkmap

import (
    "fmt"
    "testing"
)

type TcpPacket struct {
    UUID      string
    RemoteUrl string
    Name      string
}

func (t *TcpPacket) String() string {
    return fmt.Sprintf("%+v", *t)
}

func (t *TcpPacket) Key1() interface{} {
    return t.UUID
}

func (t *TcpPacket) Key2() interface{} {
    return t.RemoteUrl
}

func mmobj() *Map {
    kk := NewMap()
    kk.Set(&TcpPacket{
        UUID:      "1001",
        RemoteUrl: "127.0.0.1:8801",
        Name:      "tom",
    })

    kk.Set(&TcpPacket{
        UUID:      "1002",
        RemoteUrl: "127.0.0.1:8802",
        Name:      "jack",
    })

    kk.Set(&TcpPacket{
        UUID:      "1003",
        RemoteUrl: "127.0.0.1:8803",
        Name:      "lily",
    })

    kk.Set(&TcpPacket{
        UUID:      "1004",
        RemoteUrl: "127.0.0.1:8804",
        Name:      "lucy",
    })
    return kk
}

func TestKKmap(t *testing.T) {
    kk := mmobj()
    fmt.Println(kk.GetByKey1("1001"))
    fmt.Println(kk.GetByKey1("1002"))
    fmt.Println(kk.GetByKey2("127.0.0.1:8801"))
    fmt.Println(kk.GetByKey2("127.0.0.1:8804"))
    fmt.Println(kk.GetByKey1("1001").(*TcpPacket).String())
    fmt.Println(kk.Length())
}

func TestRange(t *testing.T) {
    kk := mmobj()
    kk.Set(&TcpPacket{
        UUID:      "1004",
        RemoteUrl: "127.0.0.1:8840",
        Name:      "poly",
    })

    fmt.Println(kk.Length())

    kk.RangeByKey1(func(k KKobject) {
        fmt.Println(k.(*TcpPacket).UUID)
    })
    kk.RangeByKey2(func(k KKobject) {
        fmt.Println(k.(*TcpPacket).RemoteUrl)
    })
}

func TestDelete(t *testing.T) {
    kk := mmobj()
    kk.Delete(&TcpPacket{
        UUID:      "1001",
        RemoteUrl: "127.0.0.1:8801",
        Name:      "",
    })

    kk.Delete(&TcpPacket{
        UUID:      "1002",
        RemoteUrl: "127.0.0.1:8822",
        Name:      "",
    })

    // kk.DeleteByKey2("127.0.0.1:8803")
    kk.DeleteByKey1("1003")

    fmt.Println(kk.Length())
    fmt.Println(kk)

}
