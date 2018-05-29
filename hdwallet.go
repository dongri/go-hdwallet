package hdwallet

import (
	"crypto/hmac"
	"crypto/sha512"
)

// var
var (
	MasterSecret   = []byte("Bitcoin seed")
	HardenedOffset = 0x80000000
	Len            = 78
)

// HDWallet ...
type HDWallet struct {
	ChainCode  []byte
	PrivateKey []byte
}

// PrivateExtendedKey ...
func (h *HDWallet) PrivateExtendedKey() {
	if len(h.PrivateKey) > 0 {
		// cs.encode(serialize(this, this.versions.private, Buffer.concat([Buffer.alloc(1, 0), this.privateKey])))
	}
}

// FromMasterSeed ...
func FromMasterSeed(seed []byte) *HDWallet {
	key := []byte(MasterSecret)
	mac := hmac.New(sha512.New, key)
	mac.Write(seed)
	I := mac.Sum(nil)
	IL := I[0:32]
	IR := I[32:]

	hdkey := new(HDWallet)
	hdkey.ChainCode = IR
	hdkey.PrivateKey = IL
	return hdkey
}
