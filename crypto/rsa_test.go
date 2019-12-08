/**
 * @Author: sybs
 * @Description:
 * @File:  rsa_test
 * @Version: 1.0.0
 * @Date: 2019-11-26 02:06
 */

package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {

	//rsa 密钥文件产生
	fmt.Println("--------------------------生成公私钥---------------------------")
	keypair := NewGenRSA(2048)

	fmt.Println(string(keypair.PrivateKey))
	fmt.Println(string(keypair.PublicKey))

	fmt.Println("------------------------签名与验签-----------------------------")
	var data = "hello world"
	signData := keypair.RsaSignWithSha256([]byte(data))

	fmt.Printf("签名信息： \n%s\n", hex.EncodeToString(signData))
	if keypair.RsaVerySignWithSha256([]byte(data), signData) {
		fmt.Printf("签名验证：%s\n", "成功")
	}

	fmt.Println("-------------------------加密解密------------------------------")
	ciphertext := keypair.RsaEncrypt([]byte(data))
	fmt.Printf("加密数据：\n%s\n", hex.EncodeToString(ciphertext))
	fmt.Printf("加密数据：\n%s\n", ToBase64(ciphertext))
	sourceData := keypair.RsaDecrypt(ciphertext)
	fmt.Printf("解密数据：\n%s\n", string(sourceData))
}


