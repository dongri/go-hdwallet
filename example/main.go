package main

import (
	"encoding/hex"
	"fmt"

	hdwallet "github.com/dongri/go-hdwallet"
)

func main() {
	hexStr := "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"
	seed, _ := hex.DecodeString(hexStr)
	masterprv := hdwallet.FromMasterSeed(seed)
	fmt.Println(masterprv)
}
