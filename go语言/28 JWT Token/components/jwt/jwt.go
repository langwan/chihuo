package jwt2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

const (
	HS256 = "HS256"
)

var alg = HS256

var Secret string

func hs256(secret, data []byte) (ret string, err error) {
	hasher := hmac.New(sha256.New, secret)
	_, err = hasher.Write(data)
	if err != nil {
		return "", err
	}
	r := hasher.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(r), nil
}

func Sign(payload interface{}) (ret string, err error) {
	h := header{
		Alg: alg,
		Typ: "JWT",
	}
	marshal, err := json.Marshal(h)
	if err != nil {
		return "", err
	}

	bh := base64.RawURLEncoding.EncodeToString(marshal)

	marshal, err = json.Marshal(payload)
	if err != nil {
		return "", err
	}

	bp := base64.RawURLEncoding.EncodeToString(marshal)

	s := fmt.Sprintf("%s.%s", bh, bp)

	ret, err = hs256([]byte(Secret), []byte(s))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s.%s", bh, bp, ret), nil
}

func Verify(token string) (err error) {
	parts := strings.Split(token, ".")
	data := strings.Join(parts[0:2], ".")
	hasher := hmac.New(sha256.New, []byte(Secret))
	_, err = hasher.Write([]byte(data))
	if err != nil {
		return err
	}
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return err
	}
	if hmac.Equal(sig, hasher.Sum(nil)) {
		return nil
	}
	return errors.New("verify is invalid")
}
