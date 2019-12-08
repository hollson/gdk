/**
 * @Author: sybs
 * @Description:RSA非对称加密算法，包含签名、验签、加密、解密
 * @File:  rsa
 * @Version: 1.0.0
 * @Date: 2019-11-26 01:27
 */

package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

//RSA长度
type RsaLength int

const (
	LEN_1024 RsaLength = 1024
	LEN_2048 RsaLength = 2048
)

type RSA struct {
	Length                RsaLength //RSA长度
	PrivateKey, PublicKey []byte    //私钥公钥
}

func NewRSA(PrvKey, PubKey []byte) *RSA {
	return &RSA{
		PrivateKey: PrvKey,
		PublicKey:  PubKey,
	}
}

//RSA公钥私钥产生
func NewGenRSA(length RsaLength) *RSA {
	rt := new(RSA)
	// 生成私钥文件
	if length == LEN_1024 || length == LEN_2048 {
		rt.Length = length
	} else {
		rt.Length = length
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, int(rt.Length))
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	rt.PrivateKey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	rt.PublicKey = pem.EncodeToMemory(block)
	return rt
}

//私钥签名
func (p *RSA) RsaSignWithSha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(p.PrivateKey)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

//公钥验签
func (p *RSA) RsaVerySignWithSha256(data, signData []byte) bool {
	block, _ := pem.Decode(p.PublicKey)
	if block == nil {
		panic(errors.New("public key error"))
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func (p *RSA) RsaEncrypt(data []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(p.PublicKey)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func (p *RSA) RsaDecrypt(ciphertext []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(p.PrivateKey)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}

func ToBase64(rsc []byte) string {
	return base64.StdEncoding.EncodeToString(rsc)
}
