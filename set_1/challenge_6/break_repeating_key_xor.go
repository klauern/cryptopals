package challenge6

import "github.com/steakknife/hamming"

// HammingDistance calculates the distance of two strings based two strings for comparison.
// It returns the overall distance as an average.
func HammingDistance(from, to []byte) int {
	// var dist = make([]int, len(from))
	var tot int
	for i := range from {
		// dist[i] = hamming.Byte(from[i], to[i])
		tot += hamming.Byte(from[i], to[i])
	}
	return tot
}

// EditDistance computes the Edit Distance (Hamming) for a Key size.
//
// For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes,
// and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
func EditDistance(keysize int, encoded []byte) int {
	keys := len(encoded) / keysize
	for i := 0; i < keys; i++ {
		// offset := i * keysize
		// end := i*keysize + keysize

		HammingDistance(encoded[i:i+keysize], encoded[i*keysize:(i*keysize)+keysize])
	}
	return -1
}
