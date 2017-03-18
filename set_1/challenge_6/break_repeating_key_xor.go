package challenge6

import "github.com/steakknife/hamming"

// HammingDistance calculates the distance of two strings based two strings for comparison.
// It returns the overall distance as an average.
func HammingDistance(from, to string) int {
	// var dist = make([]int, len(from))
	var tot int
	for i := range from {
		// dist[i] = hamming.Byte(from[i], to[i])
		tot += hamming.Byte(from[i], to[i])
	}
	return tot
}
