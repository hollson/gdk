package crypto

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

// "DES":   数据加密算法（Data Encryption Algorithm，DEA）是一种对称加密算法，很可能是使用最广泛的密钥系统，
// 特别是在保护金融数据的安全中，最初开发的DEA是嵌入硬件中的。通常，自动取款机（Automated Teller Machine，ATM）
// 都使用DEA。它出自IBM的研究工作，IBM也曾对它拥有几年的专利权，但是在1983年已到期后，处于公有范围中，允许在特定
// 条件下可以免除专利使用费而使用。1977年被美国政府正式采纳。

// Des对称加密
func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = zeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

// Des对称解密
func DesDecrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = zeroUnPadding(out)
	return string(out), nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
