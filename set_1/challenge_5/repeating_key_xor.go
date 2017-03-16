package main

import (
	"encoding/hex"
)

// RepeatingKeyXOR performs a repeating-key XOR  using the input
// 'key' and a given input string, 'input', returning the encoded string.
func RepeatingKeyXOR(key, input []byte) string {
	var inProgress []byte
	for i := 0; i < len(input); i++ {
		idx := i % len(key)
		// fmt.Printf("i is %v, with ary[idx] being %d and val %v\n", i, idx, string(key[idx]))
		ch := input[i] ^ key[idx]
		inProgress = append(inProgress, ch)
	}
	// fmt.Printf("String: %v\n", hex.EncodeToString(inProgress))
	return hex.EncodeToString(inProgress)
}

func main() {}
