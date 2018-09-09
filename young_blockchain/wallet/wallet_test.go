package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	private, public := GenerateKey()
	if private.X == nil || len(public) == 0 {
		t.Error("Generating PublicKey is error")
	}
}

func TestNewWallet(t *testing.T) {
	wallet := NewWallet()
	if len(wallet.PublicKey) == 0 {
		t.Error("NewWallet is error")
	}
}

func TestWallet_GetAddress(t *testing.T) {
	wallet := NewWallet()
	bytesAddress := wallet.GetAddress()
	if len(bytesAddress) != 36 {
		t.Error("TestCreateAddress is error")
	}
}
func TestWallet_ValidateAddress(t *testing.T) {
	isvalid := ValidateAddress("1241DqYKddenzFeaxGjNabHg6BsPVLoDBPJP")
	if isvalid == false {
		t.Error("TestWallet_ValidateAddress is error")
	}
}

func TestEncode(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey

	encPriv, encPub := Encode(privateKey, publicKey)

	fmt.Println(encPriv)
	fmt.Println(encPub)

	priv2, pub2 := Decode(encPriv, encPub)

	if !reflect.DeepEqual(privateKey, priv2) {
		fmt.Println("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, pub2) {
		fmt.Println("Public keys do not match.")
	}
}

func TestCheckSum(t *testing.T) {
	bytes := []byte{1, 2, 3}
	bytesHash := CheckSum(bytes)
	if len(bytesHash) != addressChecksumLen {
		t.Error("TestHashPubKey is error")
	}
}

func TestSha256(t *testing.T) {
	Sha256()
}

func TestRipemd160(t *testing.T) {
	Ripemd160()
}
