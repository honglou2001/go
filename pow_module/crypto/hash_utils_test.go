package crypto

import "testing"

func TestHashPubKey(t *testing.T)  {
	bytes := []byte{1,2,3}
	bytesHash := HashPubKey(bytes)
	if len(bytesHash) != 20 {
		t.Error("TestHashPubKey is error")
	}
}
