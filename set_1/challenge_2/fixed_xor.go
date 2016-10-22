package main

func FixedXor(input, inXor []byte) []byte {
	length := len(input)
	output := make([]byte, length)
	for i, in := range input {
		output[i] = in ^ inXor[i]
	}
	return output
}
