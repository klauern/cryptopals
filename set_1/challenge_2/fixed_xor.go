package main

import "encoding/hex"
import "github.com/pkg/errors"

func FixedXorString(input, inXor string) ([]byte, error) {
	if len(input) != len(inXor) {
		return []byte(""), errors.New("buffers not the same length")
	}
	inBytes, err := hex.DecodeString(input)
	if err != nil {
		return []byte(""), err
	}
	xorBytes, err := hex.DecodeString(inXor)
	if err != nil {
		return []byte(""), err
	}
	return FixedXor(inBytes, xorBytes), nil
}

func FixedXor(input, inXor []byte) []byte {
	length := len(input)
	output := make([]byte, length)
	for i, in := range input {
		output[i] = in ^ inXor[i]
	}
	return output
}
