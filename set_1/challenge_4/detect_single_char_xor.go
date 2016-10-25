package main

import "unicode"
import "github.com/klauern/cryptopals/set_1/challenge_3"

type StringCipherScore struct {
	possibilities []string
	best          string
	bestScore     int
}

func DetectSingleCharXor(lines []string) (*StringCipherScore, error) {
	var scores []*StringCipherScore
	for _, line := range lines {
		scores = append(scores, BestCipherFromString(line))
	}

	best := scores[0]
	for _, score := range scores {
		if score.bestScore > best.bestScore {
			best = score
		}
	}
	return best, nil
}

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
		best.possibilities = []string{str}
		best.best = str
		best.bestScore = score
	} else if score == best.bestScore {
		best.possibilities = append(best.possibilities, str)
	}
}
