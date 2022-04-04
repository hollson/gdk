# kkmap
A map type with two keys

# Install
```bash
go get github.com/hollson/kkmap
```

# Usage

```golang
package main

import (
    "fmt"
    "github.com/hollson/kkmap"
)

// define a kkmap object
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

func main() {
    // create a kkmap
    kk := kkmap.NewMap()
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

    fmt.Println(kk.GetByKey2("127.0.0.1:8804"))
    fmt.Println(kk.GetByKey1("1001").(*TcpPacket).String())
    fmt.Println(kk.Length())

    kk.Delete(&TcpPacket{
        UUID:      "1001",
        RemoteUrl: "127.0.0.1:8801",
        Name:      "",
    })

    kk.RangeByKey1(func(k kkmap.KKobject) {
        fmt.Println(k.(*TcpPacket).UUID)
    })
}
```
