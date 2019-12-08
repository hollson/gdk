package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var key16 = "KhBQKxfMeElLPUNS"                 //对应 AES-128
var key24 = "KhBQKxfMeElLPUNSQ66R0y9y"         //对应 AES-192
var key32 = "KhBQKxfMeElLPUNSQ66R0y9yW5kio4XM" //对应 AES-256

var txt = "DES（Data Encryption Standard）是对称加密算法"

func TestAesDecrypt16(t *testing.T) {
	result, err := AesEncrypt([]byte(txt), []byte(key16))
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, []byte(key16))
	if err != nil {
		panic(err)
	}
	if string(origData) != txt {
		panic("err")
	}
}

func TestAesDecrypt24(t *testing.T) {
	result, err := AesEncrypt([]byte(txt), []byte(key24))
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, []byte(key24))
	if err != nil {
		panic(err)
	}
	if string(origData) != txt {
		panic("err")
	}
}

func TestAesDecrypt32(t *testing.T) {
	result, err := AesEncrypt([]byte(txt), []byte(key32))
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, []byte(key32))
	if err != nil {
		panic(err)
	}
	if string(origData) != txt {
		panic("err")
	}
}

func TestAesDecryptStr16(t *testing.T) {
	result := AesEncryptSimple(txt, key16)
	fmt.Println(result)
	origData := AesDecryptSimple(result, key16)
	if origData != txt {
		panic("err")
	}
}

func TestAesDecryptStr24(t *testing.T) {
	result := AesEncryptSimple(txt, key24)
	fmt.Println(result)
	origData := AesDecryptSimple(result, key24)
	if origData != txt {
		panic("err")
	}
}

func TestAesDecryptStr32(t *testing.T) {
	result := AesEncryptSimple(txt, key32)
	fmt.Println(result)
	origData := AesDecryptSimple(result, key32)
	if origData != txt {
		panic("err")
	}
}

//=================DES=======================

//func TestDesEncrypt16(t *testing.T) {
//	result, err := DesEncrypt(txt, []byte(key16))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(result)
//	origData, err := DesDecrypt(result, []byte(key16))
//	if err != nil {
//		panic(err)
//	}
//	if string(origData) != txt {
//		panic("err")
//	}
//}