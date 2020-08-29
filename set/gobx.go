//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-08
// @ Version: 1.0.0
//
// gob 即golang binary,是go对象的二进制存储数据结构
//-------------------------------------------------------------------------------------

package set

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string
}

func Test() {
	p := Person{"Tom"}
	if err := Encode("/tmp/test.gob", p); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Encode：", "ok")

	if tar, err := Decode("/tmp/test.gob"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Decode：", tar.Name)
	}
}

// gob编码
func Encode(path string, tar interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf) //gob编码器

	if err := enc.Encode(tar); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

// gob解码：Decode参数须是具体类型，不能是interface，无法封装
func Decode(path string) (tar *Person, err error) {
	file, err := os.Open(path)
	defer file.Close()

	dec := gob.NewDecoder(file) //gob解码器
	if err = dec.Decode(&tar); err != nil {
		return nil, err
	}
	return tar, nil
}
