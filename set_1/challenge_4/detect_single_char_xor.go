package main

import "unicode"
import "github.com/klauern/cryptopals/set_1/challenge_3"

// StringCipherScore represents a list of scores for a given
// slice of possibilities
type StringCipherScore struct {
	best      string
	bestScore int
}

func worker(lineIn <-chan string, resultOut chan<- *StringCipherScore) {
	for line := range lineIn {
		best := BestCipherFromString(line)
		resultOut <- best
	}
}

// DetectSingleCharXor will take a slice of string and find the one among it
// that contains a decipherable decoding using one Xor operation against
// a character
func DetectSingleCharXor(lines []string) (*StringCipherScore, []*StringCipherScore, error) {
	var scores []*StringCipherScore

	linesCh := make(chan string, 50)
	resultsCh := make(chan *StringCipherScore, 50)

	// create workers for processing scores
	for w := 1; w <= 5; w++ {
		go worker(linesCh, resultsCh)
	}

	go func(lines []string) {
		for _, line := range lines {
			linesCh <- line
		}
		close(linesCh)
	}(lines)

	best := &StringCipherScore{}
	for range lines {
		score := <-resultsCh
		scores = append(scores, score)
	}
	for _, score := range scores {
		if score.bestScore > best.bestScore {
			best = score
		}
	}
	return best, scores, nil
}

// BestCipherFromString will send on the *StringCipherScore channel, the best
// possible cipher decoding from a given string.
func BestCipherFromString(line string) *StringCipherScore {
	best := &StringCipherScore{}
	for _, r16 := range unicode.ASCII_Hex_Digit.R16 {
		for c := r16.Lo; c <= r16.Hi; c += r16.Stride {
			best.addCipher(rune(c), []byte(line))
		}
	}
	for _, r32 := range unicode.ASCII_Hex_Digit.R32 {
		for c := r32.Lo; c <= r32.Hi; c += r32.Stride {
			best.addCipher(rune(c), []byte(line))
		}
	}
	return best
}

func (best *StringCipherScore) addCipher(c rune, line []byte) {
	str, score := challenge_3.ScoreCipher(c, line)
	if score > best.bestScore {
		best.best = str
		best.bestScore = score
	}
}
