// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package month

import (
    "fmt"
    "regexp"
    "testing"
    "time"
)

func TestInstance(t *testing.T) {
    fmt.Println(Current())
    fmt.Println(Month(202101))
    fmt.Println(New(2021, 12))
    fmt.Println(FromTime(time.Now()))
    fmt.Println(FromDot(201801))
    fmt.Println(FromTick(3))
    fmt.Println(Min())
    fmt.Println(Max())
}

func TestMethod(t *testing.T) {
    m := FromDot(201801)
    fmt.Println(m.Year())
    fmt.Println(m.Month())
    fmt.Println(m.Add(-2))
    fmt.Println(m.Add(13))
    fmt.Println(m.Prev())
    fmt.Println(m.Next())
    fmt.Println(m.Tick())
    fmt.Println(m.Diff(Month(201812)))
    fmt.Println(m.Diff(Month(201711)))
    fmt.Println(m.Quarter())
    fmt.Println(m.String())
    fmt.Println(m.Format("YY年MM月"))
    fmt.Println(m.Format("yyyy年mm月"))
    fmt.Println(m.Format("mm月yyyy年"))
    fmt.Println(m.Format("yyyy-mm yyyy-mm"))
}

func TestRange(t *testing.T) {
    Range(Month(202002), Month(202103), func(m Month) {
        fmt.Printf("%d - %d\n", m.Year(), m.Month())
    })
    fmt.Println(Span(Month(202002), Month(202103)))
}

func TestReg(t *testing.T) {
    text := "Abc a7c MFC 8ca. 你好！ Golang/"
    reg := regexp.MustCompile(`(Ab|a7)c`)
    fmt.Printf("%q\n", reg.ReplaceAllString(text, `${1}ccc`))

    text = "yy-mm"
    reg = regexp.MustCompile(`(yyyy|yy)`)
    fmt.Printf("%q\n", reg.ReplaceAllLiteralString(text, `1234`))
}
