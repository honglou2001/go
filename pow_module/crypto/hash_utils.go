package crypto

import (
	"crypto/sha256"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/ripemd160"
)

/*Hash public key*/
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		beego.Error("HashPubKey fail,", err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}
