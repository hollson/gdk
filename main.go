/**
 * @Author: sybs
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2019-11-26 10:39
 */

//$ cat hex.go
package main

// 编码问题
import (
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"fmt"
)

func Something() {
	//func EncodeToString(src []byte) string 编码byte字节为16进制字符串
	src := []byte("hello")
	fmt.Println(src)                     //[104 101 108 108 111]
	encodeStr := hex.EncodeToString(src) //68656c6c6f 16进制转换
	fmt.Println(encodeStr)

	//func Encode(dst, src []byte) int
	//func EncodedLen(n int) int
	Welcome := []byte("Gopher!")
	Wdest := make([]byte, hex.EncodedLen(len(Welcome)))
	num := hex.Encode(Wdest, Welcome)
	fmt.Println(Wdest, num) //num=14

	//func DecodeString(s string) ([]byte, error)  解码16进制的字符串为byte类型
	decodeStr, _ := hex.DecodeString(encodeStr)
	fmt.Println(string(decodeStr))

	//func DecodedLen(x int) int  x个byte解码后的长度，一般是x/2
	//func Decode(dst, src []byte) (int, error) 将byte类型的src解码到byte类型的dst中，并且返回dst的长度
	test := []byte("48656c6c6f20476f7068657221")
	dest := make([]byte, hex.DecodedLen(len(test))) //定义一个切片
	num, err := hex.Decode(dest, test)              //转换16进制字符串为byte[]类型，返回切片长度
	if err != nil {
		return
	}
	fmt.Println(num, dest[:num], string(dest), len(dest), cap(dest)) // print 13

	//func Dump(data []byte) string         //返回给定字符串以及字符串相对应的hex dump文件 效果相当于linux命令行下的"hexdump -C filename"
	content := []byte("Go is an open source programming language.")
	fmt.Println(hex.Dump(content))
	logrus.Infoln("hello")
}

func main() {
	Something()
}

//$ go run hex.go
//
//[104 101 108 108 111]
//68656c6c6f
//[52 55 54 102 55 48 54 56 54 53 55 50 50 49] 14
//hello
//13 [72 101 108 108 111 32 71 111 112 104 101 114 33] Hello Gopher! 13 13
//00000000  47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|
//00000010  75 72 63 65 20 70 72 6f  67 72 61 6d 6d 69 6e 67  |urce programming|
//00000020  20 6c 61 6e 67 75 61 67  65 2e                    | language.|
