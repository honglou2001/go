package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/astaxie/beego"
)

type Wallet struct {
	PublicKey string  			//公钥
	PrivateKey string 			//私钥
	Address string              //地址
}

var (
	runMode  string
	randKey  string
	randSign string
	prk      *ecdsa.PrivateKey
	puk      ecdsa.PublicKey
	curve   elliptic.Curve
)
func (wallet *Wallet) GeneratePublicKey() {
	var err error
	curve  = elliptic.P256()
	//创建密匙对
	prk, err = ecdsa.GenerateKey(curve,  rand.Reader)
	if err != nil {
		beego.Error("Crypt init fail,", err, " need = ", curve.Params().BitSize)
		return
	}
	puk = prk.PublicKey
}

func Encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func Decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

