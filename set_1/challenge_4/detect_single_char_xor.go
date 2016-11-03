package main

import "unicode"
import "github.com/klauern/cryptopals/set_1/challenge_3"

// StringCipherScore represents a list of scores for a given
// slice of possibilities
type StringCipherScore struct {
	possibilities []string
	best          string
	bestScore     int
}

// DetectSingleCharXor will take a slice of string and find the one among it
// that contains a decipherable decoding using one Xor operation against
// a character
func DetectSingleCharXor(lines []string) (*StringCipherScore, error) {
	var scores []*StringCipherScore
	chanScores := make(chan *StringCipherScore, 5)
	go func() {
		for _, line := range lines {
			BestCipherFromString(line, chanScores)
		}
	}()
	for i := 0; i < len(lines); i++ {
		select {
		case score := <-chanScores:
			scores = append(scores, score)
		}
	}
	best := scores[0]
	for _, score := range scores {
		if score.bestScore > best.bestScore {
			best = score
		}
	}
	return best, nil
}

// BestCipherFromString will send on the *StringCipherScore channel, the best
// possible cipher decoding from a given string.
func BestCipherFromString(line string, ch chan<- *StringCipherScore) {
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
	ch <- best
}

func (best *StringCipherScore) addCipher(c rune, line []byte) {
	str, score := challenge_3.ScoreCipher(c, line)
	if score > best.bestScore {
		best.possibilities = []string{str}
		best.best = str
		best.bestScore = score
	} else if score == best.bestScore {
		best.possibilities = append(best.possibilities, str)
	}
}
