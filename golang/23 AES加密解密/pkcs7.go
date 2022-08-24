package main

import (
	"bytes"
	"errors"
)

func pkcs7Pad(src []byte, blockSize int) (dest []byte, err error) {
	if blockSize <= 0 {
		return nil, errors.New("block size is 0")
	} else if src == nil || len(src) == 0 {
		return nil, errors.New("src is nil")
	}
	n := blockSize - (len(src) % blockSize)
	pb := make([]byte, len(src)+n)
	copy(pb, src)
	copy(pb[len(src):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func pkcs7Unpad(src []byte, blockSize int) (dest []byte, err error) {

	if blockSize <= 0 {
		return nil, errors.New("block size is 0")
	} else if len(src)%blockSize != 0 {
		return nil, errors.New("src length error")
	} else if src == nil || len(src) == 0 {
		return nil, errors.New("src is nil")
	}

	c := src[len(src)-1]

	padLength := int(c)

	if padLength == 0 || padLength > len(src) {
		return nil, errors.New("pad length error")
	}

	for i := 0; i < padLength; i++ {
		if src[len(src)-padLength+i] != c {
			return nil, errors.New("pad content error")
		}
	}

	return src[:len(src)-padLength], nil

}
