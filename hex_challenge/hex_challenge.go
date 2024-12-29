package hex_challenge

import (
	"encoding/hex"
	"fmt"
)

func HexChallenge() {
	HexToBase64()
}

func HexToBase64() {
	trial := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	dst := make([]byte, hex.DecodedLen(len(trial)))
	n, err := hex.Decode(dst, trial)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", dst[:n])
}