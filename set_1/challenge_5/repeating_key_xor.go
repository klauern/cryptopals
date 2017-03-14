package main

import (
	"fmt"
)

// RepeatingKeyXOR performs a repeating-key XOR  using the input
// 'key' and a given input string, 'input', returning the encoded string.
func RepeatingKeyXOR(key, input []byte) string {
	var inProgress []byte
	for i := 0; i < len(input); i++ {
		ch := key[i%len(key)] ^ input[i]
		inProgress = append(inProgress, ch)
	}
	fmt.Printf("%v", inProgress)
	return string(inProgress)
}

func main() {}
