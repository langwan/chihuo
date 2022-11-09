package main

import (
	"fmt"
	"github.com/dgryski/dgoogauth"
	"github.com/skip2/go-qrcode"
	"testing"
)

func TestGenQr(t *testing.T) {

	otp := dgoogauth.OTPConfig{
		Secret:      gAuthSecret,
		HotpCounter: 0,
	}

	url := otp.ProvisionURIWithIssuer("chihuo", "Company")
	fmt.Printf(url)

	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qr.png")

	if err != nil {
		t.Error(err)
	}

}

func TestVerify(t *testing.T) {
	otp := dgoogauth.OTPConfig{
		Secret:      gAuthSecret,
		HotpCounter: 0,
	}

	ok, err := otp.Authenticate("111111")
	if err != nil {
		t.Error(err)
		return
	}

	if !ok {
		t.Fail()
	}

}
