package main

// HammingDistance calculates the distance of two strings based on
// an assumed KEYSIZE value and two strings for comparison.  It returns
// the overall distance as an average.
func HammingDistance(keysize int, from, to string) int {
	// take the first keysize of bytes against the second keysize of bytes
	var dist = make([]byte, len(from))
	for i := range from {
		dist[i] = from[i] - to[i]
	}
	return 0
}
