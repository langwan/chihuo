package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func Encrypt(key, src []byte) (data []byte, err error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	} else if len(src) == 0 {
		return nil, errors.New("src is empty")
	}

	plaintext, err := pkcs7Pad(src, block.BlockSize())

	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	
	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func Decrypt(key, src []byte) (data []byte, err error) {

	if len(src) < aes.BlockSize {
		return nil, errors.New("data length error")
	}

	iv := src[:aes.BlockSize]
	ciphertext := src[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(ciphertext, ciphertext)
	ciphertext, err = pkcs7Unpad(ciphertext, aes.BlockSize)

	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
