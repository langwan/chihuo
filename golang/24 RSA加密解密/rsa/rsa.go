package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Encrypt(key *rsa.PublicKey, src []byte) (data []byte, err error) {
	h := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(h, rand.Reader, key, src, nil)

	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func Decrypt(key *rsa.PrivateKey, src []byte) (data []byte, err error) {
	h := sha256.New()
	oaep, err := rsa.DecryptOAEP(h, rand.Reader, key, src, nil)
	if err != nil {
		return nil, err
	}
	return oaep, nil
}

func Sign(key *rsa.PrivateKey, src []byte) (sign []byte, err error) {
	h := crypto.SHA256
	hn := h.New()
	hn.Write(src)
	sum := hn.Sum(nil)
	return rsa.SignPSS(rand.Reader, key, h, sum, nil)
}

func Verify(key *rsa.PublicKey, sign, src []byte) (err error) {
	h := crypto.SHA256
	hn := h.New()
	hn.Write(src)
	sum := hn.Sum(nil)
	return rsa.VerifyPSS(key, h, sum, sign, nil)
}
