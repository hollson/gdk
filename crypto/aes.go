package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//ref : http://www.blogfshare.com/aes-rijndael.html

// "AES":   高级加密标准（英语：Advanced Encryption Standard，缩写：AES），在密码学中又称Rijndael加密法，
// 			是美国联邦政府采用的一种区块加密标准。这个标准用来替代原先的DES，已经被多方分析且广为全世界所使用。经过五年的
// 			甄选流程，高级加密标准由美国国家标准与技术研究院（NIST）于2001年11月26日发布于FIPS PUB 197，并在2002年5月
// 			26日成为有效的标准。2006年，高级加密标准已然成为对称密钥加密中最流行的算法之一。

// Aes对称加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// Aes对称解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

// Aes对称加密（简单版本）
func AesEncryptSimple(plainText string, key string) string {
	ptxt, keyBts := []byte(plainText), []byte(key) //转成字节数组

	block, _ := aes.NewCipher(keyBts) //分组秘钥(cipher:密码/暗号)
	blockSize := block.BlockSize()
	ptxt = pKCS5Padding(ptxt, blockSize) //补全码

	tmp := make([]byte, len(ptxt))                            //密文容器
	cbcE := cipher.NewCBCEncrypter(block, keyBts[:blockSize]) //CBC加密器
	cbcE.CryptBlocks(tmp, ptxt)                               //加密

	return base64.StdEncoding.EncodeToString(tmp)
}

// Aes对称解密（简单版本）
func AesDecryptSimple(cipherText string, key string) string {
	ctxt, _ := base64.StdEncoding.DecodeString(cipherText)
	keyBts := []byte(key)

	block, _ := aes.NewCipher(keyBts)                         //分组秘钥
	blockSize := block.BlockSize()                            //秘钥块长度
	cbcD := cipher.NewCBCDecrypter(block, keyBts[:blockSize]) //CBC解密器

	tmp := make([]byte, len(ctxt))
	cbcD.CryptBlocks(tmp, ctxt)

	tmp = pKCS5UnPadding(tmp) //去补全码
	return string(tmp)
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
