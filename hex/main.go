package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	stringToBytes()
	bytesToHexBytes()
	bytesToHexString()
	decodeHexBytes()
	decodeHexString()
}

func stringToBytes() {
	msgBytes := []byte("hello world")
	fmt.Println(string(msgBytes))
}

func bytesToHexBytes() {
	msgBytes := []byte("hello world")
	hexBytes := make([]byte, hex.EncodedLen(len(msgBytes)))
	hex.Encode(hexBytes, msgBytes)
	fmt.Println(string(hexBytes))
}

func bytesToHexString() {
	msgBytes := []byte("hello world")
	hexString := hex.EncodeToString(msgBytes)
	fmt.Println(hexString)
}

func decodeHexBytes() {
	hexBytes := []byte("68656c6c6f20776f726c64")
	msgBytes := make([]byte, hex.EncodedLen(len(hexBytes)))
	hex.Decode(msgBytes, hexBytes)
	fmt.Println(string(msgBytes))
}

func decodeHexString() {
	hexString := "68656c6c6f20776f726c64"
	msgBytes, _ := hex.DecodeString(hexString)
	fmt.Println(string(msgBytes))
}
