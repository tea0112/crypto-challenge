package hex_challenge

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexChallenge() {
	hexStr := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Printf("%s", HexToBase64(hexStr))
}

func HexToBase64(hexStr []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(hexStr)))
	n, err := hex.Decode(dst, hexStr)
	if err != nil {
		panic(err)
	}

	base64Bytes := make([]byte, base64.StdEncoding.EncodedLen(len(dst[:n])))
	base64.StdEncoding.Encode(base64Bytes, dst)

	return base64Bytes
}
