package main

import (
	appRsa "app/rsa"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
)

func main() {

	priKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		return
	}

	pubKey := priKey.PublicKey

	src := "0123456789012345678901234567890123456789"

	encrypt, err := appRsa.Encrypt(&pubKey, []byte(src))
	if err != nil {
		fmt.Printf("encrypt err : %v", err)
		return
	}

	decrypt, err := appRsa.Decrypt(priKey, encrypt)
	if err != nil {
		fmt.Printf("decrypt err : %v", err)
		return
	}
	fmt.Printf("解密数据: %s\n", decrypt)

	sign, err := appRsa.Sign(priKey, []byte(src))
	if err != nil {
		fmt.Printf("sign err : %v", err)
		return
	}

	fmt.Printf("签名为: %s\n", hex.EncodeToString(sign))

	err = appRsa.Verify(&pubKey, sign, []byte(src))
	if err != nil {
		fmt.Printf("verify err : %v", err)
		return
	} else {
		fmt.Println("校验签名成功")
	}

}
